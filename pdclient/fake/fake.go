package fake

import "net/http"

type ClientDoer struct {
	Response   *http.Response
	Error      error
	SpyRequest http.Request
}

func (s *ClientDoer) Do(req *http.Request) (resp *http.Response, err error) {
	s.SpyRequest = *req
	return s.Response, s.Error
}
