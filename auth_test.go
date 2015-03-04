package pezdispenser_test

import (
	"github.com/go-martini/martini"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/pivotalservices/pezdispenser"
)

var _ = Describe("Authentication", func() {
	Describe("InitAuth", func() {
		var m *ClassicMartini
		BeforeEach(func() {
			m = martini.Classic()
		})

		Context("when InitAuth is passed a classic martini", func() {
			BeforeEach(func() {
				InitAuth(m)
			})

			It("Should setup the authentication middleware", func() {
				Ω(true).Should(BeTrue())
			})
		})
	})
})
