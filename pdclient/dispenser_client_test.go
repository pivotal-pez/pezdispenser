package pdclient_test

import (
	"bytes"
	"io/ioutil"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/pivotal-pez/pezdispenser/pdclient"
	"github.com/pivotal-pez/pezdispenser/pdclient/fake"
)

var _ = Describe("PDClient struct", func() {

	Describe("given a NewClient func", func() {
		Context("when called with a valid api-key and http.Client", func() {
			controlKey := "random-api-key"
			controlURL := "api.random.io"
			var pdclient *PDClient
			BeforeEach(func() {
				pdclient = NewClient(controlKey, controlURL, new(fake.ClientDoer))
			})
			It("then it should return a properly initialized pdclient", func() {
				Ω(pdclient.APIKey).Should(Equal(controlKey))
				Ω(pdclient.URL).Should(Equal(controlURL))
			})
		})
	})

	Describe("given a PostLease() method", func() {
		Context("when called with valid arguments", func() {
			controlKey := "random-api-key"
			controlURL := "api.random.io"
			controlLeaseID := "fakelease"
			controlInventoryID := "fakeinventoryid"
			controlSkuID := "fakesku"
			controlResponseBody := "{}"
			var fakeClient *fake.ClientDoer
			var pdclient *PDClient
			BeforeEach(func() {
				fakeClient = &fake.ClientDoer{
					Response: &http.Response{
						Body: ioutil.NopCloser(bytes.NewBufferString(controlResponseBody)),
					},
				}
				pdclient = NewClient(controlKey, controlURL, fakeClient)
				pdclient.PostLease(controlLeaseID, controlInventoryID, controlSkuID)
			})
			It("then it should create a valid request object", func() {
				body, _ := ioutil.ReadAll(fakeClient.SpyRequest.Body)
				Ω(body).Should(ContainSubstring(controlLeaseID))
				Ω(body).Should(ContainSubstring(controlInventoryID))
				Ω(body).Should(ContainSubstring(controlSkuID))
			})
			XIt("then it should receive the task object from the rest endpoint, parse and return it", func() {
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
