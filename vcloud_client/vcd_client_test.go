package vcloud_client_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/pivotal-pez/pezdispenser/vcloud_client"
)

var _ = Describe("VCloud Client", func() {
	Describe("VCDAuth", func() {
		It("should do nothing yet", func() {
			vcdClient := NewVCDAuth()
			token := vcdClient.GetToken()
			Ω(token).ShouldNot(BeEmpty())
		})
	})
})
