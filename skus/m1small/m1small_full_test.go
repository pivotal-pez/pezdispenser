package m1small_test

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
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
				}, 3, 1).Should(ContainSubstring(taskmanager.AgentTaskStatusFailed))
			})
		})
	})
})
