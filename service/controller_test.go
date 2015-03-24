package pezdispenser_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Random Controller", func() {
	Context("when called with some random arguments", func() {
		BeforeEach(func() {
		})

		It("Should setup do something great", func() {
			Ω(true).Should(BeTrue())
		})
	})
})
