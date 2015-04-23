package pezauth_test

import (
	"log"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/pivotalservices/pezauth/service"
)

var _ = Describe("NewMeController", func() {
	Context("calling controller", func() {
		var (
			render     *mockRenderer
			testLogger = log.New(os.Stdout, "testLogger", 0)
		)
		setGetUserInfo("pivotal.io", "jcalabrese@pivotal.io")

		BeforeEach(func() {
			render = new(mockRenderer)
		})

		It("should return a user object to the renderer", func() {
			tokens := &mockTokens{}
			controlResponse := Response{Payload: GetUserInfo(tokens)}
			var meGet MeGetHandler = NewMeController().Get().(MeGetHandler)
			meGet(testLogger, render, tokens)
			Ω(render.StatusCode).Should(Equal(200))
			Ω(render.ResponseObject).Should(Equal(controlResponse))
		})
	})
})
