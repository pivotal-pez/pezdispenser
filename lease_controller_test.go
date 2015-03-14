package pezdispenser_test

import (
	"github.com/go-martini/martini"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/pivotalservices/pezdispenser"
)

var _ = Describe("LeaseController", func() {
	Describe("LeaseTypeController", func() {
		var controller Controller

		BeforeEach(func() {
			controller = NewLeaseController(APIVersion1, Type)
		})

		Context("calling Get()", func() {
			It("Should panic", func() {
				Ω(func() {
					controller.Get()
				}).Should(Panic())
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

				It("Should not panic", func() {
					Ω(func() {
						fnc(martini.Params{TypeGUID: "something"})
					}).ShouldNot(Panic())
				})
			})
		})

		Context("calling Delete()", func() {
			It("Should panic", func() {
				Ω(func() {
					x := controller.Delete().(func(martini.Params) string)
					_ = x
				}).Should(Panic())
			})
		})
	})

	Describe("LeaseTypeController", func() {
		var controller Controller

		BeforeEach(func() {
			controller = NewLeaseController(APIVersion1, Item)
		})

		Context("calling Get()", func() {
			It("Should panic", func() {
				Ω(func() {
					controller.Get()
				}).Should(Panic())
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

				Context("with valid arguments", func() {
					controlRes := "something"
					args := martini.Params{ItemGUID: controlRes}

					It("Should not panic", func() {
						Ω(func() {
							fnc(args)
						}).ShouldNot(Panic())
					})

					It("Should return a string", func() {
						res := fnc(args)
						Ω(res).Should(Equal(controlRes))
					})
				})
			})
		})

		Context("calling Delete()", func() {
			It("Should return a function of the correct format and not panic", func() {
				Ω(func() {
					fnc := controller.Delete().(func(martini.Params) string)
					Ω(fnc).ShouldNot(BeNil())
				}).ShouldNot(Panic())
			})

			Context("calling the function returned from Delete()", func() {
				var fnc func(martini.Params) string

				BeforeEach(func() {
					fnc = controller.Delete().(func(martini.Params) string)
				})
				Context("with valid arguments", func() {
					It("Should not panic", func() {
						Ω(func() {
							fnc(martini.Params{ItemGUID: "something"})
						}).ShouldNot(Panic())
					})
				})
			})
		})
	})
})
