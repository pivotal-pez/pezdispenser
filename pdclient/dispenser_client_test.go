package pdclient_test

import (
	"bytes"
	"fmt"
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
			controlResID := "560ede8bfccecc0072000001"
			controlResTS := int64(1443815051336165844)
			controlResExpires := int64(0)
			controlResStatus := "complete"
			controlResProfile := "longpoll_queue"
			controlResCaller := "m1.small"
			controlResponseBody := fmt.Sprintf(`{
				"id": "%s","timestamp": %d,"expires": %d,"status": "%s","profile": "%s","caller_name": "%s",
				"meta_data": {}
			}`, controlResID, controlResTS, controlResExpires, controlResStatus, controlResProfile, controlResCaller)
			var (
				leaseCreateResponse LeaseCreateResponseBody
				res                 *http.Response
				err                 error
				fakeClient          *fake.ClientDoer
				pdclient            *PDClient
			)
			BeforeEach(func() {
				fakeClient = &fake.ClientDoer{
					Response: &http.Response{
						StatusCode: http.StatusOK,
						Body:       ioutil.NopCloser(bytes.NewBufferString(controlResponseBody)),
					},
				}
				pdclient = NewClient(controlKey, controlURL, fakeClient)
				leaseCreateResponse, res, err = pdclient.PostLease(controlLeaseID, controlInventoryID, controlSkuID, 14)
			})
			It("then it should create a valid request object", func() {
				body, _ := ioutil.ReadAll(fakeClient.SpyRequest.Body)
				Ω(body).Should(ContainSubstring(controlLeaseID))
				Ω(body).Should(ContainSubstring(controlInventoryID))
				Ω(body).Should(ContainSubstring(controlSkuID))
			})
			It("then it should receive the task object from the rest endpoint, parse and return it", func() {
				Ω(err).ShouldNot(HaveOccurred())
				Ω(res.StatusCode).Should(Equal(http.StatusOK))
				Ω(leaseCreateResponse.ID).Should(Equal(controlResID))
				Ω(leaseCreateResponse.Timestamp).Should(Equal(controlResTS))
				Ω(leaseCreateResponse.Expires).Should(Equal(controlResExpires))
				Ω(leaseCreateResponse.Status).Should(Equal(controlResStatus))
				Ω(leaseCreateResponse.Profile).Should(Equal(controlResProfile))
				Ω(leaseCreateResponse.CallerName).Should(Equal(controlResCaller))
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
