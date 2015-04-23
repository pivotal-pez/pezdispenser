package pezdispenser_test

import (
	"github.com/go-martini/martini"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/pivotalservices/pezdispenser/service"
)

var _ = Describe("Routes", func() {
	Describe("InitRoutes", func() {
		var m *martini.ClassicMartini
		BeforeEach(func() {
			m = martini.Classic()
		})

		Context("when InitRoutes is passed a classic martini", func() {

			It("Should not result in panic", func() {
				Ω(func() {
					InitRoutes(m, "testurl.org")
				}).ShouldNot(Panic())
			})
		})
	})
})
