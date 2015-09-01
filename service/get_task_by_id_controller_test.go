package pezdispenser_test

import (
	"log"
	"net/http"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pivotal-pez/pezdispenser/fakes"
	. "github.com/pivotal-pez/pezdispenser/service"
	"github.com/pivotal-pez/pezdispenser/service/integrations"
	"github.com/pivotal-pez/pezdispenser/taskmanager"
)

var _ = Describe("GetTaskByIdController()", func() {
	Context("when the handler response is called with a valid ID params value", func() {
		var (
			fakeURI              = "mongodb://c39642c7-xxxx-xxxx-xxxx-db67a3bbc98f:xxxx4b827xxxx4393dcxxxx3533xxxx@1.1.1.1:27017/70ef645b-xxxx-xxxx-xxxx-94d5b0e5107f"
			handler              func(params martini.Params, log *log.Logger, r render.Render, t integrations.Collection)
			controlID            = "myvalidid"
			renderer             *fakes.FakeRenderer
			controlResponseValue = taskmanager.Task{
				Status: "rockin and rollin",
				MetaData: map[string]interface{}{
					"some": "stuff",
				},
			}
		)

		BeforeEach(func() {
			taskCollection := SetupDB(fakes.FakeNewCollectionDialer(controlResponseValue), fakeURI, TaskCollectionName)
			handler = GetTaskByIDController().(func(params martini.Params, log *log.Logger, r render.Render, t integrations.Collection))
			renderer = new(fakes.FakeRenderer)
			handler(martini.Params{"id": controlID}, fakes.MockLogger, renderer, taskCollection)
		})

		It("should return the RedactedTasktask object w/ a 200 statusCode", func() {
			Ω(renderer.SpyStatus).Should(Equal(http.StatusOK))
			Ω(func() {
				_ = renderer.SpyValue.(*taskmanager.Task)
			}).Should(Panic())
			Ω(renderer.SpyValue.(taskmanager.RedactedTask)).Should(Equal(controlResponseValue.GetRedactedVersion()))
		})
	})
})
