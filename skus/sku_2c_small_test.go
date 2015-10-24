package skus_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pivotal-pez/pezdispenser/fakes"
	"github.com/pivotal-pez/pezdispenser/skurepo"
	. "github.com/pivotal-pez/pezdispenser/skus"
	"github.com/pivotal-pez/pezdispenser/taskmanager"
	"github.com/pivotal-pez/pezdispenser/vcloudclient"
)

var _ = Describe("Sku2CSmall", func() {
	Describe(".PollForTasks()", func() {
		Context("when the outsourced task is found to be in a success state", func() {

			It("should expire the current task and set status to success", func() {
				s, spyTask, _ := fakes.MakeFakeSku2CSmall(vcloudclient.TaskStatusSuccess)
				s.PollForTasks()
				Ω(spyTask.Status).Should(Equal(vcloudclient.TaskStatusSuccess))
				Ω(spyTask.Expires).Should(Equal(taskmanager.ExpiredTask))
			})
		})

		Context("when task action value is un-deploy and status is success", func() {
			var (
				sku          *Sku2CSmall
				spyTask      *taskmanager.Task
				spySavedTask *taskmanager.Task
			)
			BeforeEach(func() {
				sku, spyTask, spySavedTask = fakes.MakeFakeSku2CSmall(vcloudclient.TaskStatusSuccess)
				spyTask.PrivateMetaData = map[string]interface{}{
					taskmanager.TaskActionMetaName: TaskActionUnDeploy,
					VCDUsernameField:               "fakeuser",
					VCDPasswordField:               "fakepass",
					VCDTemplateNameField:           "PaaSSlot-10",
					VCDAppIDField:                  "myapp-id",
					VCDBaseURIField:                "vcd_base_uri.com",
				}
			})
			It("should move on to deploy a new 2c vapp", func() {
				Ω(spySavedTask).ShouldNot(Equal(spyTask))
				sku.PollForTasks()
				Ω(spySavedTask).ShouldNot(BeNil())
				Ω(spySavedTask).Should(Equal(spyTask))
			})
		})

		Context("when task action value is deploy and status is success", func() {
			var (
				sku          *Sku2CSmall
				spyTask      *taskmanager.Task
				spySavedTask *taskmanager.Task
			)
			BeforeEach(func() {
				sku, spyTask, spySavedTask = fakes.MakeFakeSku2CSmall(vcloudclient.TaskStatusSuccess)
				spyTask.PrivateMetaData = map[string]interface{}{
					taskmanager.TaskActionMetaName: TaskActionDeploy,
					VCDUsernameField:               "fakeuser",
					VCDPasswordField:               "fakepass",
					VCDTemplateNameField:           "PaaSSlot-10",
					VCDAppIDField:                  "myapp-id",
					VCDBaseURIField:                "vcd_base_uri.com",
				}
			})
			It("should show inventory as available", func() {
				Ω(spySavedTask).ShouldNot(Equal(spyTask))
				sku.PollForTasks()
				Ω(spySavedTask.Status).Should(Equal(taskmanager.TaskStatusAvailable))
			})
		})

		Context("when task action value is self-destruct and status is success", func() {
			var (
				sku          *Sku2CSmall
				spyTask      *taskmanager.Task
				spySavedTask *taskmanager.Task
			)
			BeforeEach(func() {
				sku, spyTask, spySavedTask = fakes.MakeFakeSku2CSmall(vcloudclient.TaskStatusSuccess)
				spyTask.PrivateMetaData = map[string]interface{}{
					taskmanager.TaskActionMetaName: TaskActionSelfDestruct,
					VCDUsernameField:               "fakeuser",
					VCDPasswordField:               "fakepass",
					VCDTemplateNameField:           "PaaSSlot-10",
					VCDAppIDField:                  "myapp-id",
					VCDBaseURIField:                "vcd_base_uri.com",
				}
			})
			It("should have kicked off a restock task", func() {
				sku.PollForTasks()
				Ω(spySavedTask.PrivateMetaData[taskmanager.TaskActionMetaName]).Should(Equal(TaskActionUnDeploy))
			})
			It("should have cleaned up the state of the sku object", func() {
				sku.PollForTasks()
				Ω(sku.ProcurementMeta).Should(BeNil())
			})
		})

		Context("when the outsourced task is found to be in a not yet done state", func() {
			It("should update status and move on", func() {
				s, spyTask, _ := fakes.MakeFakeSku2CSmall(vcloudclient.TaskStatusRunning)
				controlStatus := spyTask.Status
				controlExpires := spyTask.Expires
				s.PollForTasks()
				Ω(spyTask.Status).ShouldNot(Equal(controlStatus))
				Ω(spyTask.Expires).Should(Equal(controlExpires))
				Ω(spyTask.Expires).ShouldNot(Equal(taskmanager.ExpiredTask))
			})
		})
		Context("when the outsourced task is found to be in a failed state", func() {
			Context("with a status of error", func() {
				s, spyTask, _ := fakes.MakeFakeSku2CSmall(vcloudclient.TaskStatusError)

				It("should expire the task and set status error", func() {
					s.PollForTasks()
					Ω(spyTask.Status).Should(Equal(vcloudclient.TaskStatusError))
					Ω(spyTask.Expires).Should(Equal(taskmanager.ExpiredTask))
				})
			})
			Context("with a status of aborted", func() {
				s, spyTask, _ := fakes.MakeFakeSku2CSmall(vcloudclient.TaskStatusAborted)

				It("should expire the task and set status of aborted", func() {
					s.PollForTasks()
					Ω(spyTask.Status).Should(Equal(vcloudclient.TaskStatusAborted))
					Ω(spyTask.Expires).Should(Equal(taskmanager.ExpiredTask))
				})
			})
			Context("with a status of canceled", func() {
				s, spyTask, _ := fakes.MakeFakeSku2CSmall(vcloudclient.TaskStatusCanceled)

				It("should expire the task and set status canceled", func() {
					s.PollForTasks()
					Ω(spyTask.Status).Should(Equal(vcloudclient.TaskStatusCanceled))
					Ω(spyTask.Expires).Should(Equal(taskmanager.ExpiredTask))
				})
			})
		})
	})

	Describe(".New()", func() {
		Context("when called with valid arguments", func() {
			It("should return an initialized Sku interface object", func() {
				controlMeta := map[string]interface{}{
					"base_uri": "random.com",
				}
				s := new(Sku2CSmall)
				s.ProcurementMeta = map[string]interface{}{
					LeaseExpiresFieldName: time.Now().UnixNano(),
				}
				sku := s.New(new(fakes.FakeTaskManager), controlMeta)
				skuCast := sku.(*Sku2CSmall)
				Ω(skuCast.ProcurementMeta).Should(Equal(controlMeta))
			})
		})
	})

	Describe(".Procurement()", func() {
		Context("when called with valid metadata", func() {
			var (
				task               *taskmanager.Task
				fakeTaskManager    = new(fakes.FakeTaskManager)
				controlInventoryID = "random-guid"
			)
			BeforeEach(func() {
				s := new(Sku2CSmall)
				s.ProcurementMeta = map[string]interface{}{
					LeaseExpiresFieldName: time.Now().UnixNano(),
					InventoryIDFieldName:  controlInventoryID,
					VCDTemplateNameField:  "pcfaas-slot-10",
				}
				fakeTaskManager.SpyTaskSaved = new(taskmanager.Task)
				sku := s.New(fakeTaskManager, s.ProcurementMeta)
				skuCast := sku.(*Sku2CSmall)
				task = skuCast.Procurement()
			})

			It("should return a status complete", func() {
				Ω(task.Status).Should(Equal(StatusComplete))
			})

			It("then it should return meta data with a creds field", func() {
				Ω(task.MetaData).ShouldNot(BeEmpty())
				Ω(task.MetaData[CredentialsFieldName]).ShouldNot(BeNil())
			})

			It("should create a self-destruct lease task", func() {
				Ω(fakeTaskManager.SpyTaskSaved).ShouldNot(BeNil())
				Ω(fakeTaskManager.SpyTaskSaved.GetPrivateMeta(taskmanager.TaskActionMetaName)).Should(Equal(TaskActionSelfDestruct))
				Ω(fakeTaskManager.SpyTaskSaved.GetPrivateMeta(InventoryIDFieldName)).Should(Equal(controlInventoryID))
			})
		})
	})

	Describe(".ReStock()", func() {
		Context("when called with valid metadata", func() {
			var sku skurepo.Sku
			controlTaskHref := "myfakehref"
			BeforeEach(func() {
				fakeClient := new(fakes.FakeVCDClient)
				fakeClient.FakeVAppTemplateRecord = new(vcloudclient.VAppTemplateRecord)
				fakeClient.FakeVAppTemplateRecord.Href = "fakehref"
				fakeClient.FakeVAppTemplateRecord.Vdc = "fakevdchref"
				fakeClient.FakeVApp = new(vcloudclient.VApp)
				fakeClient.FakeVApp.Tasks = vcloudclient.TasksElem{}
				fakeClient.FakeVApp.Tasks.Task = vcloudclient.TaskElem{}
				fakeClient.FakeVApp.Tasks.Task.Href = controlTaskHref

				sku = &Sku2CSmall{
					Client:          fakeClient,
					TaskManager:     new(fakes.FakeTaskManager),
					ProcurementMeta: make(map[string]interface{}),
				}
			})
			It("should return a status indicating the current status", func() {
				task := sku.ReStock()
				Ω(task.Status).Should(Equal(StatusOutsourced))
			})
		})
	})
})
