package pezdispenser_test

import (
	"encoding/json"

	"github.com/go-martini/martini"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/pivotal-pez/pezdispenser/service"
)

var _ = Describe("LeaseController", func() {
	Describe("LeaseListController", func() {
		var controller Controller
		var fnc func() string

		BeforeEach(func() {
			controller = NewLeaseController(APIVersion1, List)
		})

		It("calling Get()", func() {
			Ω(func() {
				fnc = controller.Get().(func() string)
			}).ShouldNot(Panic())
		})

		Context("calling the function returned from Get w/ valid args", func() {
			It("should return a valid response object", func() {
				res := fnc()
				resObj := &ResponseMessage{}
				json.Unmarshal([]byte(res), resObj)
				Ω(len(resObj.Body)).Should(BeNumerically(">", 0))
			})
		})
	})
	Describe("LeaseTypeController", func() {
		var controller Controller

		BeforeEach(func() {
			controller = NewLeaseController(APIVersion1, Type)
		})

		Context("calling Get()", func() {
			It("Should return a function of the correct format and not panic", func() {
				Ω(func() {
					fnc := controller.Get().(func(martini.Params) string)
					Ω(fnc).ShouldNot(BeNil())
				}).ShouldNot(Panic())
			})

			Context("calling the function returned from Get()", func() {
				var fnc func(martini.Params) string

				BeforeEach(func() {
					fnc = controller.Get().(func(martini.Params) string)
				})

				Context("with valid arguments", func() {
					controlRes := "something"
					args := martini.Params{TypeGUID: controlRes}

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
					args := martini.Params{TypeGUID: controlRes}

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
					x := controller.Delete().(func(martini.Params) string)
					_ = x
				}).Should(Panic())
			})
		})
	})

	Describe("LeaseItemController", func() {
		var controller Controller

		BeforeEach(func() {
			controller = NewLeaseController(APIVersion1, Item)
		})

		Context("calling Get()", func() {
			It("Should return a function of the correct format and not panic", func() {
				Ω(func() {
					fnc := controller.Get().(func(martini.Params) string)
					Ω(fnc).ShouldNot(BeNil())
				}).ShouldNot(Panic())
			})

			Context("calling the function returned from Get()", func() {
				var fnc func(martini.Params) string

				BeforeEach(func() {
					fnc = controller.Get().(func(martini.Params) string)
				})

				Context("with valid arguments", func() {
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
	})
})
