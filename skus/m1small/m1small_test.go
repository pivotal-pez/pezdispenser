package m1small_test

import (
	"os"
	"sync"
	"sync/atomic"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pivotal-pez/pezdispenser/fakes"
	"github.com/pivotal-pez/pezdispenser/innkeeperclient"
	"github.com/pivotal-pez/pezdispenser/innkeeperclient/fake"
	. "github.com/pivotal-pez/pezdispenser/skus/m1small"
	"github.com/pivotal-pez/pezdispenser/taskmanager"
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
	Describe("given .StartPoller() method", func() {
		Context("when a status polling call to innkeeper yields a 'complete' status", func() {
			var (
				fakeInnKeeperClient *fakeinnkeeperclient.IKClient
				controlInventoryID  = "id-something"
				controlRequestID    = "request-test-id"
				spyTask             *taskmanager.Task
			)
			BeforeEach(func() {
				spyTask = new(taskmanager.Task)
				fakeInnKeeperClient = new(fakeinnkeeperclient.IKClient)
				fakeInnKeeperClient.StatusCallCountForComplete = int64(0)
				fakeInnKeeperClient.FakeData = make([]innkeeperclient.RequestData, 1)
				fakeInnKeeperClient.FakeData[0] = innkeeperclient.RequestData{
					RequestID: controlRequestID,
				}
				s := new(SkuM1Small)
				s.ProcurementMeta = map[string]interface{}{
					"lease_expires": time.Now().UnixNano(),
					"inventory_id":  controlInventoryID,
				}
				fakeTaskManager := &fakes.FakeTaskManager{
					SpyTaskSaved: new(taskmanager.Task),
				}
				spyTask.Protect(fakeTaskManager, sync.RWMutex{})
				sku := s.New(fakeTaskManager, s.ProcurementMeta)
				skuCast := sku.(*SkuM1Small)
				skuCast.Client = fakeInnKeeperClient
				fakeInnKeeperClient.SpyStatusCallCount = new(int64)
				skuCast.StartPoller(controlRequestID, spyTask)
			})
			It("then it should update the original lease task with 'complete' status", func() {
				Ω(spyTask.Status).Should(Equal(taskmanager.AgentTaskStatusComplete))
			})
			It("then it should update the original lease task with expired state", func() {
				Ω(spyTask.Expires).Should(Equal(taskmanager.ExpiredTask))
			})
			It("then it should update the original lease task metadata with innkeeper info", func() {
				controlStatusResponse := innkeeperclient.GetStatusResponse{
					Status: taskmanager.AgentTaskStatusComplete,
					Data:   innkeeperclient.Data{Status: taskmanager.AgentTaskStatusComplete},
				}
				Ω(spyTask.MetaData[GetStatusInfoMetaName].(innkeeperclient.GetStatusResponse)).Should(Equal(controlStatusResponse))
			})
		})
	})
	Describe("given .Procurement() method", func() {
		Context("when a provision call to innkeeper is complete", func() {
			var (
				fakeInnKeeperClient *fakeinnkeeperclient.IKClient
				controlInventoryID  = "id-something"
				controlRequestID    = "request-test-id"
			)
			BeforeEach(func() {
				fakeInnKeeperClient = new(fakeinnkeeperclient.IKClient)
				fakeInnKeeperClient.FakeData = make([]innkeeperclient.RequestData, 1)
				fakeInnKeeperClient.FakeData[0] = innkeeperclient.RequestData{
					RequestID: controlRequestID,
				}
				s := new(SkuM1Small)
				s.ProcurementMeta = map[string]interface{}{
					"lease_expires": time.Now().UnixNano(),
					"inventory_id":  controlInventoryID,
				}
				fakeTaskManager := &fakes.FakeTaskManager{
					SpyTaskSaved: new(taskmanager.Task),
				}
				sku := s.New(fakeTaskManager, s.ProcurementMeta)
				skuCast := sku.(*SkuM1Small)
				skuCast.Client = fakeInnKeeperClient
				fakeInnKeeperClient.SpyStatusCallCount = new(int64)
				skuCast.Procurement()
			})

			It("then it should start the innkeeper status poller", func() {
				Eventually(func() int64 {
					return atomic.LoadInt64(fakeInnKeeperClient.SpyStatusCallCount)
				}).Should(BeNumerically(">", int64(0)))
			})

		})
		Context("when called with valid input", func() {
			var (
				task               *taskmanager.Task
				fakeTaskManager    *fakes.FakeTaskManager
				controlInventoryID = "random-guid"
				controlMessage     = "hi there"
				controlStatus      = "complete"
				controlPHInfo      = &innkeeperclient.ProvisionHostResponse{
					Data: []innkeeperclient.RequestData{
						innkeeperclient.RequestData{},
					},
					Message: controlMessage,
					Status:  controlStatus,
				}
				controlClient *fakeinnkeeperclient.IKClient
				skuCast       *SkuM1Small
			)
			BeforeEach(func() {
				controlClient = &fakeinnkeeperclient.IKClient{
					FakeMessage:        []string{controlMessage},
					FakeStatus:         []string{controlStatus},
					SpyStatusCallCount: new(int64),
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
				}).Should(Equal(controlStatus))
			})

		})
	})
})
