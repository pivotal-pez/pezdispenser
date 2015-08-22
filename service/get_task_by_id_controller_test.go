package pezdispenser_test

import (
	"bytes"
	"log"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/pivotal-pez/pezdispenser/service"
)

type fakeRenderer struct {
	render.Render
	SpyStatus int
	SpyValue  interface{}
}

func (s *fakeRenderer) JSON(status int, v interface{}) {
	s.SpyStatus = status
	s.SpyValue = v
}

var _ = Describe("GetTaskByIdController()", func() {
	Context("when the handler response is called with a valid ID params value", func() {
		var (
			handler              func(params martini.Params, log *log.Logger, r render.Render)
			controlID            = "myvalidid"
			renderer             *fakeRenderer
			controlResponseValue = map[string]string{"taskID": controlID}
		)

		BeforeEach(func() {
			handler = GetTaskByIdController("taskURI").(func(params martini.Params, log *log.Logger, r render.Render))
			renderer = new(fakeRenderer)
			var buf bytes.Buffer
			logger := log.New(&buf, "logger: ", log.Lshortfile)
			handler(martini.Params{"id": controlID}, logger, renderer)
		})

		It("should return the task object w/ a 200 statusCode", func() {
			Ω(renderer.SpyStatus).Should(Equal(200))
			Ω(renderer.SpyValue).Should(Equal(controlResponseValue))
		})
	})
})
