package pezdispenser_test

import (
	"github.com/go-martini/martini"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/pivotalservices/pezdispenser"
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

			It("Should not panic", func() {
				Ω(func() {
					fnc(martini.Params{ItemGUID: "something"})
				}).ShouldNot(Panic())
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

			It("Should not panic", func() {
				Ω(func() {
					fnc(martini.Params{ItemGUID: "something"})
				}).ShouldNot(Panic())
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
