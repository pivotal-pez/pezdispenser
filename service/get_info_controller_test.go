package pezdispenser_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/pivotal-pez/pezdispenser/service"
)

var _ = Describe("GetInfoController()", func() {
	Context("when we call the returned Handler object", func() {
		It("should give us the info message", func() {
			handler := GetInfoController().(func() string)
			Ω(handler()).ShouldNot(BeEmpty())
		})
	})
})
