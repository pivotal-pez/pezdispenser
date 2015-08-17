package vcloud_client_test

import (
	"errors"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/pivotal-pez/pezdispenser/vcloud_client"
)

var _ = Describe("VCloud Client", func() {
	Describe("VCDClient", func() {

		Describe(".Auth()", func() {
			var (
				vcdClient    *VCDClient
				controlToken = "xxxxxxxxxxxxxxxxxedw8d8sdb9sdb9sdbsd9sdbsdb"
			)

			Context("given valid user/pass/uri", func() {
				BeforeEach(func() {
					client := new(fakeHttpClient)
					client.Response = new(http.Response)
					client.Response.StatusCode = AuthSuccessStatusCode
					client.Response.Header = http.Header{}
					client.Response.Header.Set(VCloudTokenHeaderName, controlToken)
					vcdClient = NewVCDClient(client)
				})

				It("should set us a valid auth token", func() {
					err := vcdClient.Auth("", "", "")
					token := vcdClient.Token
					Ω(err).ShouldNot(HaveOccurred())
					Ω(token).ShouldNot(BeEmpty())
					Ω(token).Should(Equal(controlToken))
				})
			})

			Context("given the api does not authenticate against our credentials", func() {
				BeforeEach(func() {
					client := new(fakeHttpClient)
					client.Response = new(http.Response)
					client.Response.StatusCode = (AuthSuccessStatusCode + 201)
					client.Response.Header = http.Header{}
					client.Response.Header.Set(VCloudTokenHeaderName, controlToken)
					vcdClient = NewVCDClient(client)
				})

				It("should return the proper error", func() {
					err := vcdClient.Auth("", "", "")
					Ω(err).Should(HaveOccurred())
					Ω(err).Should(Equal(ErrAuthFailure))
				})

				It("should not set a token", func() {
					vcdClient.Auth("", "", "")
					token := vcdClient.Token
					Ω(token).Should(BeEmpty())
				})
			})

			Context("given an authentication call returns error", func() {
				BeforeEach(func() {
					client := new(fakeHttpClient)
					client.Err = errors.New("random connection error")
					client.Response = new(http.Response)
					client.Response.StatusCode = (AuthSuccessStatusCode + 300)
					vcdClient = NewVCDClient(client)
				})

				It("should pass through the error from the client connection", func() {
					err := vcdClient.Auth("", "", "")
					token := vcdClient.Token
					Ω(err).Should(HaveOccurred())
					Ω(err).ShouldNot(Equal(ErrAuthFailure))
					Ω(token).Should(BeEmpty())
				})
			})
		})
	})
})
