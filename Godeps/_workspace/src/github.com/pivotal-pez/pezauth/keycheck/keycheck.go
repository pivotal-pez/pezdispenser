package keycheck

import "net/http"

var (
	HeaderKeyName = "X-API-KEY"
)

//New - a new keychecker targetted at the given url
func New(url string) KeyChecker {
	return &keyCheckV1{
		target: url,
		client: new(http.Client),
	}
}

type (
	//ClientDoer - an interface for a http Client.Do
	ClientDoer interface {
		Do(req *http.Request) (resp *http.Response, err error)
	}
	//KeyChecker - an interface for something that can check a key
	KeyChecker interface {
		SetClient(ClientDoer)
		Check(string) (res *http.Response, err error)
	}
	keyCheckV1 struct {
		target string
		client ClientDoer
	}
)

//Check - checks given key against the targetted validator endpoint
func (s *keyCheckV1) Check(key string) (res *http.Response, err error) {
	req, err := http.NewRequest("GET", s.target, nil)
	req.Header.Add(HeaderKeyName, key)
	res, err = s.client.Do(req)
	return
}

//SetClient
func (s *keyCheckV1) SetClient(client ClientDoer) {
	s.client = client
}
