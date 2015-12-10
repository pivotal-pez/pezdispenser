package skurepo_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pivotal-pez/pezdispenser/fakes"
	. "github.com/pivotal-pez/pezdispenser/skurepo"
	"github.com/pivotal-pez/pezdispenser/vcloudclient"
)

var _ = Describe("repo", func() {
	var (
		controlSkuKey = "myRegisteredSku"
		mySku, _, _   = fakes.MakeFakeSku(vcloudclient.TaskStatusSuccess)
	)
	Describe("given: a Register() method", func() {
		Context("when: passed a name and a Sku interface", func() {
			BeforeEach(func() {
				Register(controlSkuKey, mySku)
			})
			AfterEach(func() {
				Repo = make(map[string]Sku)
			})
			It("then: it should add the given Sku under the given name in the registry", func() {
				registry := GetRegistry()
				Ω(registry).ShouldNot(BeEmpty())
				Ω(registry[controlSkuKey]).Should(Equal(mySku))
			})
		})
	})

	Describe("given: a GetRegistry() method", func() {
		Context("when: called without any registrerd Skus", func() {
			var registry map[string]Sku
			BeforeEach(func() {
				registry = GetRegistry()
			})
			AfterEach(func() {
				Repo = make(map[string]Sku)
			})
			It("then: it should return an empty map of Sku interfaces", func() {
				Ω(registry).Should(BeEmpty())
				Ω(registry).ShouldNot(BeNil())
			})
		})
		Context("when: called containing registrerd Skus", func() {
			var registry map[string]Sku
			BeforeEach(func() {
				Register(controlSkuKey, mySku)
				registry = GetRegistry()
			})
			AfterEach(func() {
				Repo = make(map[string]Sku)
			})
			It("then: it should return the map of registered Sku interfaces", func() {
				Ω(registry).ShouldNot(BeEmpty())
				Ω(registry[controlSkuKey]).Should(Equal(mySku))
			})
		})
	})
})
