package vcloud_client

import "net/http"

type (
	//VCDAuth - vcd authentication object
	VCDClient struct {
		Token  string
		client httpClientDoer
	}

	httpClientDoer interface {
		Do(req *http.Request) (resp *http.Response, err error)
	}
)
