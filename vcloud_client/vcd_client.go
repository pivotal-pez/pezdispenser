package vcloud_client

import (
	"crypto/tls"
	"net/http"
)

//NewVCDClient - constructs a new VCDClient object with given client
func NewVCDClient(client httpClientDoer) *VCDClient {
	return &VCDClient{
		client: client,
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

//Auth - authenticates against the vcd api and sets a token
func (s *VCDClient) Auth(username, password, uri string) (err error) {
	var (
		req   *http.Request
		resp  *http.Response
		token string
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
