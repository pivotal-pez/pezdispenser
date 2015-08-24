package skus_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/pivotal-pez/pezdispenser/skus"
)

var _ = Describe("Sku2CSmall", func() {
	Describe(".Procurement()", func() {
		Context("when called with valid metadata", func() {
			It("should return a status complete", func() {
				s := Sku2CSmall{}
				status, _ := s.Procurement(make(map[string]interface{}))
				Ω(status).Should(Equal(StatusComplete))
			})
		})
	})

	Describe(".ReStock()", func() {
		Context("when called with valid metadata", func() {
			It("should return a status complete", func() {
				s := Sku2CSmall{}
				status, _ := s.ReStock(make(map[string]interface{}))
				Ω(status).Should(Equal(StatusComplete))
			})
		})
	})
})
