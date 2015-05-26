package pezauth_test

import (
	"fmt"
	"log"
	"net/http"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/pivotalservices/pezauth/service"
)

var _ = Describe("NewValidateV1", func() {
	var (
		username   = "testuser@pivotal.io"
		guid       = "51ceee02-1111-1111-1111-ac4fe04ad2fd"
		keyhash    = fmt.Sprintf("%s:%s", username, guid)
		render     *mockRenderer
		testLogger = log.New(os.Stdout, "testLogger", 0)
		req        = &http.Request{}
		header     = &http.Header{}
		keyGen     = getKeygen(false, keyhash, false)
	)
	setGetUserInfo("pivotal.io", username)

	Context("calling controller without headers", func() {
		BeforeEach(func() {
			render = new(mockRenderer)
		})

		AfterEach(func() {
			req = &http.Request{}
			header = &http.Header{}
		})

		It("should return an error status and responseBody", func() {
			controlResponse := &Response{ErrorMsg: ErrInvalidKeyFormatMsg}
			var validGet ValidateGetHandler = NewValidateV1(keyGen).Get().(ValidateGetHandler)
			validGet(testLogger, render, req)
			Ω(render.StatusCode).Should(Equal(FailureStatus))
			Ω(render.ResponseObject).Should(Equal(*controlResponse))
		})
	})

	Context("calling controller with an in-valid key format", func() {
		BeforeEach(func() {
			header.Add(HeaderKeyName, "invalid key value")
			req.Header = *header
			render = new(mockRenderer)
		})

		AfterEach(func() {
			req = &http.Request{}
			header = &http.Header{}
		})

		It("should return an error status and responseBody", func() {
			controlResponse := &Response{ErrorMsg: ErrInvalidKeyFormatMsg}
			var validGet ValidateGetHandler = NewValidateV1(keyGen).Get().(ValidateGetHandler)
			validGet(testLogger, render, req)
			Ω(render.StatusCode).Should(Equal(FailureStatus))
			Ω(render.ResponseObject).Should(Equal(*controlResponse))
		})
	})

	Context("calling controller with a valid key", func() {
		BeforeEach(func() {
			header.Add(HeaderKeyName, guid)
			req.Header = *header
			render = new(mockRenderer)
		})

		AfterEach(func() {
			req = &http.Request{}
			header = &http.Header{}
		})

		It("should return an success status and valid responseBody", func() {
			var validGet ValidateGetHandler = NewValidateV1(keyGen).Get().(ValidateGetHandler)
			_, controlPayload, _ := keyGen.GetByKey(guid)
			validGet(testLogger, render, req)
			Ω(render.StatusCode).Should(Equal(SuccessStatus))
			Ω(render.ResponseObject.(Response).Payload).Should(Equal(controlPayload))
			Ω(render.ResponseObject.(Response).APIKey).Should(Equal(guid))
		})
	})
})
