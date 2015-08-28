package skus_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pivotal-pez/pezdispenser/fakes"
	. "github.com/pivotal-pez/pezdispenser/skus"
	"github.com/pivotal-pez/pezdispenser/vcloudclient"
)

var _ = Describe("Sku2CSmall", func() {
	Describe(".PollForTasks()", func() {
		Context("when called", func() {
			It("should do nothing right now", func() {
				s := new(Sku2CSmall)
				s.Name = SkuName2CSmall
				s.TaskManager = new(fakes.FakeTaskManager)
				s.PollForTasks()
				Ω(true).Should(BeTrue())
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
				sku := s.New(new(fakes.FakeTaskManager), controlMeta)
				skuCast := sku.(*Sku2CSmall)
				Ω(skuCast.ProcurementMeta).Should(Equal(controlMeta))
			})
		})
	})

	Describe(".Procurement()", func() {
		Context("when called with valid metadata", func() {
			It("should return a status complete", func() {
				s := new(Sku2CSmall)
				status, _ := s.Procurement()
				Ω(status).Should(Equal(StatusComplete))
			})
		})
	})

	Describe(".ReStock()", func() {
		Context("when called with valid metadata", func() {
			var sku Sku
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
				status, meta := sku.ReStock()
				Ω(status).Should(Equal(StatusOutsourced))
				Ω(meta["vcd_task_element_href"]).Should(Equal(controlTaskHref))
				Ω(meta["task_action"]).Should(Equal(TaskActionUnDeploy))
				Ω(meta["subtask_id"]).ShouldNot(BeEmpty())
			})
		})
	})
})
