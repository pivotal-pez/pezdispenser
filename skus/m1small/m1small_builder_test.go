package m1small_test

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/pivotal-pez/pezdispenser/skus/m1small"
)

var _ = Describe("Skum1small", func() {
	BeforeEach(func() {
		os.Setenv("VCAP_APPLICATION", GetDefaultVCAPApplicationString())
		os.Setenv("VCAP_SERVICES", GetDefaultVCAPServicesString())
	})
	Describe("given .New() method", func() {
		Context("when called", func() {
			It("then it should not panic", func() {
				Ω(func() {
					new(SkuM1SmallBuilder).New(nil, nil)
				}).ShouldNot(Panic())
			})
		})
	})
})
