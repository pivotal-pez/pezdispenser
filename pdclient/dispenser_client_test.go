package pdclient_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/pivotal-pez/pezdispenser/pdclient"
	"github.com/pivotal-pez/pezdispenser/pdclient/fake"
)

var _ = Describe("PDClient struct", func() {

	Describe("given a NewClient func", func() {
		Context("when called with a valid api-key and http.Client", func() {
			controlKey := "random-api-key"
			var pdclient *PDClient
			BeforeEach(func() {
				pdclient = NewClient(controlKey, new(fake.ClientDoer))
			})
			It("then it should return a properly initialized pdclient", func() {
				Ω(pdclient.APIKey).Should(Equal(controlKey))
			})
		})
	})

	XDescribe("given a PostLease() method call", func() {
		Context("when called with a valid taskguid", func() {
			It("then it should receive the task object from the rest endpoint, parse and return it", func() {
				Ω(true).Should(Equal(false))
			})
		})
	})

	XDescribe("given a GetTask(id) method call", func() {
		Context("when called with a valid taskguid", func() {
			It("then it should receive the task object from the rest endpoint, parse and return it", func() {
				Ω(true).Should(Equal(false))
			})
		})
	})

	XDescribe("given a DeleteLease() method call", func() {
		Context("when called with a valid taskguid", func() {
			It("then it should receive the task object from the rest endpoint, parse and return it", func() {
				Ω(true).Should(Equal(false))
			})
		})
	})
})
