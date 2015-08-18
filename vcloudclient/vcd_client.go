package vcloudclient

import (
	"crypto/tls"
	"fmt"
	"net/http"
)

//NewVCDClient - constructs a new VCDClient object with given client
func NewVCDClient(client httpClientDoer, baseURI string) *VCDClient {
	return &VCDClient{
		BaseURI: baseURI,
		client:  client,
	}
}

//DefaultClient - grabs a default http client for us to use for api calls
func DefaultClient() (client *http.Client) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client = &http.Client{Transport: tr}
	return
}

//AuthDecorate - decorate a given request object with a auth token header
func (s *VCDClient) AuthDecorate(req *http.Request) (err error) {

	if s.Token == "" {
		err = ErrNoTokenToApply

	} else {

		if req.Header == nil {
			req.Header = http.Header{}
		}
		req.Header.Set(VCloudTokenHeaderName, s.Token)
	}
	return
}

func (s *VCDClient) getAbsoluteURIFromPath(uriPath string) (absoluteURI string) {
	absoluteURI = fmt.Sprintf("%s%s", s.BaseURI, uriPath)
	return
}

//Auth - authenticates against the vcd api and sets a token
func (s *VCDClient) Auth(username, password string) (err error) {
	var (
		req   *http.Request
		resp  *http.Response
		token string
		uri   = s.getAbsoluteURIFromPath(VCDAuthURIPath)
	)
	defer func() {
		s.Token = token
	}()

	if req, err = http.NewRequest("POST", uri, nil); err == nil {
		req.SetBasicAuth(username, password)
		req.Header.Set("Accept", "application/*+xml;version=5.5")

		if resp, err = s.client.Do(req); err == nil && resp.StatusCode == AuthSuccessStatusCode {
			token = resp.Header.Get(VCloudTokenHeaderName)

		} else if err == nil && resp.StatusCode != AuthSuccessStatusCode {
			err = ErrAuthFailure
		}
	}
	return
}
