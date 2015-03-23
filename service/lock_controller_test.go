package pezdispenser_test

import (
	"github.com/go-martini/martini"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/pivotalservices/pezdispenser/service"
)

var _ = Describe("LockController", func() {
	var controller Controller

	BeforeEach(func() {
		controller = NewLockController(APIVersion1)
	})

	Context("calling Get()", func() {
		It("Should return a function of the correct format and not panic", func() {
			Ω(func() {
				fnc := controller.Get().(func(martini.Params) string)
				Ω(fnc).ShouldNot(BeNil())
			}).ShouldNot(Panic())
		})

		Context("calling the function returned from Post()", func() {
			var fnc func(martini.Params) string

			BeforeEach(func() {
				fnc = controller.Get().(func(martini.Params) string)
			})

			Context("when given valid arguments", func() {
				controlRes := "something"
				args := martini.Params{ItemGUID: controlRes}

				It("Should not panic", func() {
					Ω(func() {
						fnc(martini.Params{ItemGUID: "something"})
					}).ShouldNot(Panic())
				})

				Context("string response from controller", func() {
					var res string

					BeforeEach(func() {
						res = fnc(args)
					})

					It("Should return a valid response object", func() {
						isValidResponseMessage(res, controlRes)
					})
				})

			})
		})
	})

	Context("calling Post()", func() {
		It("Should return a function of the correct format and not panic", func() {
			Ω(func() {
				fnc := controller.Post().(func(martini.Params) string)
				Ω(fnc).ShouldNot(BeNil())
			}).ShouldNot(Panic())
		})

		Context("calling the function returned from Post()", func() {
			var fnc func(martini.Params) string

			BeforeEach(func() {
				fnc = controller.Post().(func(martini.Params) string)
			})

			Context("when given valid arguments", func() {
				controlRes := "something"
				args := martini.Params{ItemGUID: controlRes}

				It("Should not panic", func() {
					Ω(func() {
						fnc(args)
					}).ShouldNot(Panic())
				})

				Context("string response from controller", func() {
					var res string

					BeforeEach(func() {
						res = fnc(args)
					})

					It("Should return a valid response object", func() {
						isValidResponseMessage(res, controlRes)
					})
				})
			})
		})
	})

	Context("calling Delete()", func() {
		It("Should panic", func() {
			Ω(func() {
				controller.Delete()
			}).Should(Panic())
		})
	})
})
