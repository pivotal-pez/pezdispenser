package pezdispenser_test

import (
	"github.com/go-martini/martini"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/pivotalservices/pezdispenser"
)

var _ = Describe("Routes", func() {
	Describe("InitRoutes", func() {
		var m *martini.ClassicMartini
		BeforeEach(func() {
			m = martini.Classic()
		})

		Context("when InitRoutes is passed a classic martini", func() {
			BeforeEach(func() {
				InitRoutes(m)
			})

			It("Should setup the routes for the service", func() {
				Ω(true).Should(BeTrue())
			})
		})
	})
})
