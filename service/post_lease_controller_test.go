package pezdispenser_test

import (
	"bytes"
	"log"
	"net/http"
	"time"

	"github.com/martini-contrib/render"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pivotal-pez/pezdispenser/fakes"
	. "github.com/pivotal-pez/pezdispenser/service"
	"github.com/pivotal-pez/pezdispenser/taskmanager"
)

var _ = Describe("PostLeaseController()", func() {
	Context("when the handler response is called with a lease params value", func() {
		var (
			fakeURI              = "mongodb://c39642c7-xxxx-xxxx-xxxx-db67a3bbc98f:xxxx4b827xxxx4393dcxxxx3533xxxx@1.1.1.1:27017/70ef645b-xxxx-xxxx-xxxx-94d5b0e5107f"
			handler              func(log *log.Logger, r render.Render, req *http.Request)
			renderer             *fakes.FakeRenderer
			controlResponseValue = taskmanager.Task{
				Status: "rockin and rollin",
				MetaData: map[string]interface{}{
					"some": "stuff",
				},
			}
		)

		BeforeEach(func() {
			handler = PostLeaseController(fakeURI, fakes.FakeNewCollectionDialer(controlResponseValue)).(func(log *log.Logger, r render.Render, req *http.Request))
			renderer = new(fakes.FakeRenderer)
			request := new(http.Request)
			request.Body = fakes.FakeResponseBody{bytes.NewBufferString(`{
				"_id": "917397-292735-98293752935",
				"inventory_id": "kaasd9sd9-98239h23h9-99h3ba993ba9h3ab",
				"username": "someone",
				"lease_duration": 14
			}`)}
			handler(fakes.MockLogger, renderer, request)
		})

		It("should return the task object w/ a 200 statusCode", func() {
			Ω(renderer.SpyStatus).Should(Equal(http.StatusCreated))
			Ω(renderer.SpyValue.(*Lease).Task.Status).Should(Equal(TaskStatusUnavailable))
			Ω(renderer.SpyValue.(*Lease).Task.Timestamp).ShouldNot(Equal(time.Time{}))
		})
	})
})
