package pdclient

import "net/http"

type (
	PDClient struct {
		APIKey string
		client clientDoer
	}
	clientDoer interface {
		Do(req *http.Request) (resp *http.Response, err error)
	}
)
