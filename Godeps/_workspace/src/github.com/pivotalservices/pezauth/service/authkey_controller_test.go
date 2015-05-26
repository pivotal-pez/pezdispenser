package pezauth_test

import (
	"fmt"
	"log"
	"os"

	"github.com/go-martini/martini"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/pivotalservices/pezauth/service"
)

var _ = Describe("AuthKeyController", func() {
	var (
		controller Controller
		fakeGUID   = "123asdkghasdiawlkehgaweh"
		fakeUser   = "testuser@pivotal.io"
		fakeHash   = fmt.Sprintf("%s:%s", fakeUser, fakeGUID)
		domain     = "pivotal.io"
		testLogger = log.New(os.Stdout, "testLogger", 0)
	)
	setGetUserInfo(domain, fakeUser)

	Context("PUT", func() {
		Context("called successfully", func() {
			BeforeEach(func() {
				kg := getKeygen(false, fakeHash, false)
				controller = NewAuthKeyV1(kg)
			})

			It("Should NOT result in panic", func() {
				Ω(func() {
					controller.Put()
				}).ShouldNot(Panic())
			})

			Context("Handler function", func() {
				It("Should yeild a valid status code and response object", func() {
					var ph AuthPutHandler
					ph = controller.Put().(AuthPutHandler)
					render := &mockRenderer{}
					ph(martini.Params{UserParam: fakeUser}, testLogger, render, new(mockTokens))
					Ω(render.StatusCode).Should(Equal(SuccessStatus))
					Ω(render.ResponseObject.(Response).APIKey).Should(Equal(fakeGUID))
				})
			})
		})

		Context("called with failure", func() {
			BeforeEach(func() {
				kg := getKeygen(true, fakeHash, false)
				controller = NewAuthKeyV1(kg)
			})

			Context("Handler function", func() {
				It("Should yeild a error status code and error response object", func() {
					var ph AuthPutHandler
					ph = controller.Put().(AuthPutHandler)
					render := &mockRenderer{}
					ph(martini.Params{UserParam: fakeUser}, testLogger, render, new(mockTokens))
					Ω(render.StatusCode).Should(Equal(FailureStatus))
					Ω(render.ResponseObject.(Response).APIKey).Should(Equal(""))
					Ω(render.ResponseObject.(Response).ErrorMsg).ShouldNot(Equal(""))
				})
			})
		})
	})

	Context("GET", func() {
		Context("called successfully", func() {
			BeforeEach(func() {
				kg := getKeygen(false, fakeHash, false)
				controller = NewAuthKeyV1(kg)
			})

			It("Should NOT result in panic", func() {
				Ω(func() {
					controller.Get()
				}).ShouldNot(Panic())
			})

			Context("Handler function", func() {
				It("Should yeild a valid status code and response object", func() {
					var ph AuthGetHandler
					ph = controller.Get().(AuthGetHandler)
					render := &mockRenderer{}
					ph(martini.Params{UserParam: fakeUser}, testLogger, render, new(mockTokens))
					Ω(render.StatusCode).Should(Equal(SuccessStatus))
					Ω(render.ResponseObject.(Response).APIKey).Should(Equal(fakeGUID))
				})
			})
		})

		Context("called with failure", func() {
			BeforeEach(func() {
				kg := getKeygen(true, fakeHash, false)
				controller = NewAuthKeyV1(kg)
			})

			Context("Handler function", func() {
				It("Should yeild a error status code and error response object", func() {
					var ph AuthGetHandler
					ph = controller.Get().(AuthGetHandler)
					render := &mockRenderer{}
					ph(martini.Params{UserParam: fakeUser}, testLogger, render, new(mockTokens))
					Ω(render.StatusCode).Should(Equal(FailureStatus))
					Ω(render.ResponseObject.(Response).APIKey).Should(Equal(""))
					Ω(render.ResponseObject.(Response).ErrorMsg).ShouldNot(Equal(""))
				})
			})
		})
	})

	Context("DELETE", func() {
		Context("called successfully", func() {
			BeforeEach(func() {
				kg := getKeygen(false, fakeHash, false)
				controller = NewAuthKeyV1(kg)
			})

			It("Should NOT result in panic", func() {
				Ω(func() {
					controller.Delete()
				}).ShouldNot(Panic())
			})

			Context("Handler function", func() {
				It("Should yeild a valid status code and response object", func() {
					var ph AuthDeleteHandler
					ph = controller.Delete().(AuthDeleteHandler)
					render := &mockRenderer{}
					ph(martini.Params{UserParam: fakeUser}, testLogger, render, new(mockTokens))
					Ω(render.StatusCode).Should(Equal(SuccessStatus))
					Ω(render.ResponseObject.(Response).APIKey).ShouldNot(Equal(fakeGUID))
					Ω(render.ResponseObject.(Response).APIKey).Should(Equal(""))
				})
			})
		})

		Context("called with failure", func() {
			BeforeEach(func() {
				kg := getKeygen(true, fakeHash, false)
				controller = NewAuthKeyV1(kg)
			})

			Context("Handler function", func() {
				It("Should yeild a error status code and error response object", func() {
					var ph AuthDeleteHandler
					ph = controller.Delete().(AuthDeleteHandler)
					render := &mockRenderer{}
					ph(martini.Params{UserParam: fakeUser}, testLogger, render, new(mockTokens))
					Ω(render.StatusCode).Should(Equal(FailureStatus))
					Ω(render.ResponseObject.(Response).APIKey).Should(Equal(""))
					Ω(render.ResponseObject.(Response).ErrorMsg).ShouldNot(Equal(""))
				})
			})
		})
	})
})
