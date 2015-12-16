package m1small_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pivotal-pez/pezdispenser/fakes"
	"github.com/pivotal-pez/pezdispenser/innkeeperclient"
	"github.com/pivotal-pez/pezdispenser/innkeeperclient/fake"
	. "github.com/pivotal-pez/pezdispenser/skus/m1small"
	"github.com/pivotal-pez/pezdispenser/taskmanager"
	"os"
	"time"
)


var _ = Describe("Skum1small", func() {
	BeforeEach(func() {
		os.Setenv("VCAP_APPLICATION", GetDefaultVCAPApplicationString())
		os.Setenv("VCAP_SERVICES", GetDefaultVCAPServicesString())
	})
	Describe("given .GetInnkeeperClient() method", func() {
		Context("when called", func() {
			var (
				fakeTaskManager    *fakes.FakeTaskManager
				controlInventoryID = "random-guid"
				skuCast            *SkuM1Small
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
			It("should produce new innkeeperclient", func() {
				_, err := skuCast.GetInnkeeperClient()
				Ω(err).ShouldNot(HaveOccurred())
			})
		})
	})
	Describe("given .IsEnabled() method", func() {
		Context("when called with VCAP context setup", func() {
			It("should return true", func() {
				Ω(IsEnabled()).Should(Equal(true))
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
				skuCast       *SkuM1Small
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

		})
	})
})
