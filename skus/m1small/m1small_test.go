package m1small_test

import (
	"os"
	"time"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pivotal-pez/pezdispenser/fakes"
	"github.com/pivotal-pez/pezdispenser/innkeeperclient"
	"github.com/pivotal-pez/pezdispenser/innkeeperclient/fake"
	. "github.com/pivotal-pez/pezdispenser/skus/m1small"
	"github.com/pivotal-pez/pezdispenser/taskmanager"
)

var VCAP_SERVICES = `{
"user-provided": [
        {
          "name": "pezvalidator-service",
          "label": "user-provided",
          "tags": [],
          "credentials": {
            "target-url": "https://hcfdev.pezapp.io/valid-key"
          },
          "syslog_drain_url": ""
        },
        {
          "name": "innkeeper-service",
          "label": "user-provided",
          "tags": [],
          "credentials": {
            "enable": "1",
            "uri": "http://innkeeper.pivotal.io",
            "password": "SomePass",
            "user": "admin"
          },
          "syslog_drain_url": ""
        }
      ]
}
`
var VCAP_APPLICATION = `{
      "limits": {
        "mem": 1024,
        "disk": 1024,
        "fds": 16384
      },
      "application_version": "0",
      "application_name": "r",
      "application_uris": [
        "pivotal.io"
      ],
      "version": "0",
      "name": "r",
      "space_name": "z",
      "space_id": "4",
      "uris": [
        "pivotal.io"
      ],
      "users": null
		}`

var _ = Describe("Skum1small", func() {
	BeforeEach(func() {
		os.Setenv("VCAP_APPLICATION", VCAP_APPLICATION)
		os.Setenv("VCAP_SERVICES", VCAP_SERVICES)
	})
	Describe("given .GetInnkeeperClient() method", func() {
		Context("when called", func() {
			var (
				fakeTaskManager    *fakes.FakeTaskManager
				controlInventoryID = "random-guid"
				skuCast  *SkuM1Small
			)
			BeforeEach(func() {
				s := new(SkuM1Small)
				s.ProcurementMeta = map[string]interface{}{
					"lease_expires": time.Now().UnixNano(),
					"inventory_id":  controlInventoryID,
				}
				fakeTaskManager = &fakes.FakeTaskManager{
					SpyTaskSaved: new(taskmanager.Task),
				}
				sku := s.New(fakeTaskManager, s.ProcurementMeta)
				skuCast = sku.(*SkuM1Small)
			})
			It("should produce new innkeeperclient", func(){
				skuCast.GetInnkeeperClient()
			})
		})
		})	
	Describe("given .Procurement() method", func() {
		Context("when called with valid input", func() {
			var (
				task               *taskmanager.Task
				fakeTaskManager    *fakes.FakeTaskManager
				controlInventoryID = "random-guid"
				controlMessage     = "hi there"
				controlStatus      = "complete"
				controlPHInfo      = &innkeeperclient.ProvisionHostResponse{
					Message: controlMessage,
					Status:  controlStatus,
				}
				controlClient *fakeinnkeeperclient.IKClient
				skuCast  *SkuM1Small
			)
			BeforeEach(func() {
				controlClient = &fakeinnkeeperclient.IKClient{
					FakeMessage: []string{controlMessage},
					FakeStatus:  []string{controlStatus},
				}
				s := new(SkuM1Small)
				s.Client = controlClient
				s.ProcurementMeta = map[string]interface{}{
					"lease_expires": time.Now().UnixNano(),
					"inventory_id":  controlInventoryID,
				}
				fakeTaskManager = &fakes.FakeTaskManager{
					SpyTaskSaved: new(taskmanager.Task),
				}
				sku := s.New(fakeTaskManager, s.ProcurementMeta)
				skuCast = sku.(*SkuM1Small)
				task = skuCast.Procurement()
			})
			It("then it should call the innkeeper service to initiate m1small procurement process", func() {
				Ω(task.Status).Should(Equal(taskmanager.AgentTaskStatusRunning))
			})
			It("then it should call and wait for response from innkeeper client", func() {
				Eventually(func() interface{} {
					return task.Read(func(t *taskmanager.Task) interface{} {
						return t.GetPublicMeta(ProvisionHostInfoMetaName)
					})

				}).Should(Equal(controlPHInfo))
			})

			It("then it should update the exit status on the task", func() {
				Eventually(func() interface{} {
					return task.Read(func(t *taskmanager.Task) interface{} {
						return t.Status
					})
				}).Should(Equal(taskmanager.AgentTaskStatusComplete))
			})
			It("then poll should do nothing", func(){
				skuCast.PollForTasks()
			})
			It("then Restock should do nothing", func(){
				skuCast.ReStock()
			})

			It("then New should produce a new provider", func(){
				skuCast.New(fakeTaskManager, skuCast.ProcurementMeta)
			})

		})
	})
})
