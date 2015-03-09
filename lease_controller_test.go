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
			controller = NewLeaseController(ApiVersion1, Type)
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
		})

		Context("calling Delete()", func() {
			It("Should return a function of the correct format and not panic", func() {
				Ω(func() {
					fnc := controller.Delete().(func(martini.Params) string)
					Ω(fnc).ShouldNot(BeNil())
				}).ShouldNot(Panic())
			})
		})
	})

	Describe("LeaseTypeController", func() {
		var controller Controller

		BeforeEach(func() {
			controller = NewLeaseController(ApiVersion1, Item)
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
		})

		Context("calling Delete()", func() {
			It("Should return a function of the correct format and not panic", func() {
				Ω(func() {
					fnc := controller.Delete().(func(martini.Params) string)
					Ω(fnc).ShouldNot(BeNil())
				}).ShouldNot(Panic())
			})
		})
	})
})
