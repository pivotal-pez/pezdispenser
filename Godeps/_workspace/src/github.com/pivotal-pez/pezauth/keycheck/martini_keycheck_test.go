package keycheck_test

import (
	"log"
	"net/http"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/pivotal-pez/pezauth/keycheck"
)

var _ = Describe("keycheck.NewKeyCheckHandler", func() {
	var (
		testLogger         = log.New(os.Stdout, "testLogger", 0)
		testResponseWriter = &mockResponseWriter{}
		testRequest        = new(http.Request)
		handler            APIKeyCheckHandler
	)

	BeforeEach(func() {
		testResponseWriter = &mockResponseWriter{}
		testRequest = new(http.Request)
		mw := NewAPIKeyCheckMiddleware("fakeurl.org")
		handler = mw.Handler().(APIKeyCheckHandler)
	})

	Context("called without a key in the header", func() {
		It("should return error code and responsebody", func() {
			handler(testLogger, testResponseWriter, testRequest)
			Ω(testResponseWriter.StatusCode).Should(Equal(AuthFailStatus))
			Ω(testResponseWriter.Body).Should(Equal(AuthFailureResponse))
		})
	})

	Context("called with a invalid key in the header", func() {
		It("should return error code and responsebody", func() {
			header := http.Header{}
			header.Set(HeaderKeyName, "502-eeec-48d2")
			testRequest.Header = header
			handler(testLogger, testResponseWriter, testRequest)
			Ω(testResponseWriter.StatusCode).Should(Equal(AuthFailStatus))
			Ω(testResponseWriter.Body).Should(Equal(AuthFailureResponse))
		})
	})

	Context("called with valid key in the request header", func() {
		BeforeEach(func() {
			mockClient := &mockClientDoer{connectfail: false}
			mw := NewAPIKeyCheckMiddleware("fakeurl.org")
			mw.Keycheck.SetClient(mockClient)
			handler = mw.Handler().(APIKeyCheckHandler)
		})

		It("should not write anything to the response", func() {
			statusUnset := 0
			header := http.Header{}
			header.Set(HeaderKeyName, "51ceee02-eeec-48d2-a03a-ac4fe04ad2fd")
			testRequest.Header = header
			handler(testLogger, testResponseWriter, testRequest)
			Ω(testResponseWriter.StatusCode).Should(Equal(statusUnset))
			Ω(testResponseWriter.Body).Should(BeNil())
		})
	})
})
