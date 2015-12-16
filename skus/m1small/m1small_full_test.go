package m1small_test

import (
	"net/http"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
	"github.com/pivotal-pez/pezdispenser/fakes"
	"github.com/pivotal-pez/pezdispenser/service"
	"github.com/pivotal-pez/pezdispenser/skus/m1small"
	"github.com/pivotal-pez/pezdispenser/taskmanager"
)

var badInnKeeperURI = "http://DOESNOTEXISTS.pivotal.io"
var innkeeperUser = "admin"
var innkeeperPassword = "password"

var _ = Describe("Get SKU Lease", func() {
	Describe("NewLease", func() {
		BeforeEach(func() {
			os.Setenv("VCAP_APPLICATION", m1small.GetDefaultVCAPApplicationString())
		})
		Context("calling Procurement with a bad host", func() {
			var (
				db    = new(fakes.FakeCollection)
				lease *pezdispenser.Lease
			)
			BeforeEach(func() {
				os.Setenv("VCAP_SERVICES", m1small.GetVCAPServicesString(1, badInnKeeperURI, innkeeperUser, innkeeperPassword))
				m1small.Init()
				availableSKUS := pezdispenser.GetAvailableInventory(db)
				lease = pezdispenser.NewLease(db, availableSKUS)
				lease.Sku = m1small.SkuName
				lease.ProcurementMeta = make(map[string]interface{})
			})
			It("should fail the lease with an error", func() {
				Ω(lease).ShouldNot(BeNil())
				task := lease.Procurement()
				Ω(task.Status).Should(Equal(taskmanager.AgentTaskStatusRunning))
				Eventually(func() interface{} {
					return task.Read(func(t *taskmanager.Task) interface{} {
						return t.Status
					})
				}, 4, 1).Should(ContainSubstring(taskmanager.AgentTaskStatusFailed))
			})
		})
		Context("calling Procurement with a good host and auth", func() {
			var (
				db     = new(fakes.FakeCollection)
				lease  *pezdispenser.Lease
				server *ghttp.Server
			)
			BeforeEach(func() {
				server = ghttp.NewServer()
				os.Setenv("VCAP_SERVICES", m1small.GetVCAPServicesString(1, server.URL(), innkeeperUser, innkeeperPassword))
				m1small.Init()
				availableSKUS := pezdispenser.GetAvailableInventory(db)
				lease = pezdispenser.NewLease(db, availableSKUS)
				lease.Sku = m1small.SkuName
				lease.ProcurementMeta = make(map[string]interface{})
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyBasicAuth(innkeeperUser, innkeeperPassword),
						ghttp.RespondWith(http.StatusOK, `{ "status": "success", "data": [{"requestid": "28ac758e-a02c-11e5-9531-0050569b9b57"}], "message": "ok" }`),
					),
				)
			})
			AfterEach(func() {
				server.Close()
			})
			It("should get a lease with no error", func() {
				Ω(lease).ShouldNot(BeNil())
				task := lease.Procurement()
				Ω(task.Status).ShouldNot(ContainSubstring(taskmanager.AgentTaskStatusFailed))
			})
		})
		XContext("something to  do  soon", func() {
			It("nothing yet", func() {
				/*Ω(task.Status).Should(Equal(taskmanager.AgentTaskStatusRunning))
				Eventually(func() interface{} {
					return task.Read(func(t *taskmanager.Task) interface{} {
						return t.Status
					})
				}, 3, 1).Should(Equal(taskmanager.AgentTaskStatusComplete))
				phinfo := task.GetPublicMeta("phinfo").(*innkeeperclient.ProvisionHostResponse)
				Ω(phinfo.Status).Should(Equal("success"))
				Ω(phinfo.Data[0].RequestID).ShouldNot(BeNil())

				task2 := new(taskmanager.Task)
				fmt.Println(task.ID)
				if err1 := db.FindOne(task.ID.Hex(), task2); err1 == nil {
					fmt.Println("task2", task2)
				} else {
					fmt.Println(err1.Error())
				}*/
			})
		})
	})
})
