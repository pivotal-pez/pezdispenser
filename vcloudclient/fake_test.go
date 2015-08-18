package vcloudclient_test

import "net/http"

type fakeHttpClient struct {
	Err      error
	Response *http.Response
}

func (s *fakeHttpClient) Do(req *http.Request) (*http.Response, error) {
	return s.Response, s.Err
}
