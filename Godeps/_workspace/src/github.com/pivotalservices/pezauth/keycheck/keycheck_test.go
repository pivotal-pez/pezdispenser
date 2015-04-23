package keycheck_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/pivotalservices/pezauth/keycheck"
)

var _ = Describe("keycheck.New", func() {
	Describe("Check", func() {

		Context("when called with a in-valid target url", func() {
			var kc KeyChecker

			BeforeEach(func() {
				kc = New("badtargeturl.org")
				kc.SetClient(&mockClientDoer{connectfail: true})
			})

			It("should return an error", func() {
				res, err := kc.Check("")
				Ω(res).Should(BeNil())
				Ω(err).ShouldNot(BeNil())
			})
		})

		Context("when called with a valid target url", func() {
			var (
				kc         KeyChecker
				mockClient *mockClientDoer
			)

			BeforeEach(func() {
				kc = New("goodurl.org")
				mockClient = &mockClientDoer{connectfail: false}
				kc.SetClient(mockClient)
			})

			It("should make the call and get a valid response", func() {
				res, err := kc.Check("")
				Ω(res).ShouldNot(BeNil())
				Ω(err).Should(BeNil())
			})

			It("should pass the key to check in the header", func() {
				controlKey := "mykey"
				kc.Check(controlKey)
				Ω(mockClient.Request.Header.Get(HeaderKeyName)).Should(Equal(controlKey))
			})
		})
	})
})
