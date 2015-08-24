package skus_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pivotal-pez/pezdispenser/fakes"
	. "github.com/pivotal-pez/pezdispenser/skus"
	"github.com/pivotal-pez/pezdispenser/vcloudclient"
)

var _ = Describe("Sku2CSmall", func() {
	Describe(".Procurement()", func() {
		Context("when called with valid metadata", func() {
			It("should return a status complete", func() {
				s := Sku2CSmall{}
				status, _ := s.Procurement(make(map[string]interface{}))
				Ω(status).Should(Equal(StatusComplete))
			})
		})
	})

	Describe(".ReStock()", func() {
		Context("when called with valid metadata", func() {
			var sku Sku
			BeforeEach(func() {
				fakeClient := new(fakes.FakeVCDClient)
				fakeClient.FakeVAppTemplateRecord = new(vcloudclient.VAppTemplateRecord)
				fakeClient.FakeVAppTemplateRecord.Href = "fakehref"
				fakeClient.FakeVAppTemplateRecord.Vdc = "fakevdchref"
				fakeClient.FakeVApp = new(vcloudclient.VApp)
				fakeClient.FakeVApp.Tasks = vcloudclient.TasksElem{}
				fakeClient.FakeVApp.Tasks.Task = vcloudclient.TaskElem{}
				fakeClient.FakeVApp.Tasks.Task.Href = "faketaskhref"
				sku = &Sku2CSmall{
					Client: fakeClient,
				}
			})
			It("should return a status indicating that we outsourced the task to the vcd api", func() {
				status, _ := sku.ReStock(make(map[string]interface{}))
				Ω(status).Should(Equal(StatusOutsourced))
			})
		})
	})
})
