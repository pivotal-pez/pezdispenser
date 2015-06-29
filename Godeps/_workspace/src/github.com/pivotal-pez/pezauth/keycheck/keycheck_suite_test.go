package keycheck_test

import (
	"errors"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/pivotal-pez/pezauth/keycheck"

	"testing"
)

func TestPezValidator(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Pez Auth Validator Suite")
}

type mockClientDoer struct {
	connectfail bool
	Request     *http.Request
}

func (s *mockClientDoer) Do(req *http.Request) (resp *http.Response, err error) {
	s.Request = req

	if s.connectfail {
		err = errors.New("call failed error")

	} else {
		resp = new(http.Response)
		resp.StatusCode = AuthSuccessStatus
	}
	return
}

type mockResponseWriter struct {
	StatusCode int
	Body       []byte
}

func (s *mockResponseWriter) WriteHeader(i int) {
	s.StatusCode = i
}

func (s *mockResponseWriter) Header() (r http.Header) {
	return
}

func (s *mockResponseWriter) Write(x []byte) (a int, b error) {
	s.Body = x
	return
}
