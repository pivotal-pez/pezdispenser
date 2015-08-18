package vcloudclient_test

import (
	"io"
	"net/http"
)

type fakeHttpClient struct {
	Err      error
	Response *http.Response
}

func (s *fakeHttpClient) Do(req *http.Request) (*http.Response, error) {
	return s.Response, s.Err
}

type nopCloser struct {
	io.Reader
}

func (nopCloser) Close() error { return nil }
