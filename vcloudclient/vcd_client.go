package vcloudclient

import (
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
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

//DeployVApp - deploy a vapptemplate as a vapp to a given slot
func (s *VCDClient) DeployVApp(templateName, templateHref, vcdHref string) (res *http.Response, err error) {
	var req *http.Request
	uri := fmt.Sprintf("%s%s", vcdHref, vCDVAppDeploymentPath)
	requestBody := fmt.Sprintf(vAppDeploymentPayload, templateName, templateHref)

	if req, err = http.NewRequest("POST", uri, strings.NewReader(requestBody)); err == nil {
		s.AuthDecorate(req)
		req.Header.Set("Accept", "application/*+xml;version=5.5")
		req.Header.Set("Content-Type", vAppDeploymentContentType)
		res, err = s.client.Do(req)
	}
	return
}

//Auth - authenticates against the vcd api and sets a token
func (s *VCDClient) Auth(username, password string) (err error) {
	var (
		req   *http.Request
		resp  *http.Response
		token string
		uri   = s.getAbsoluteURIFromPath(vCDAuthURIPath)
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

//QueryTemplate - query the vcd api for the template we would like to use
func (s *VCDClient) QueryTemplate(templateName string) (vappTemplate *VAppTemplateRecord, err error) {
	var (
		req *http.Request
	)
	URI := s.getAbsoluteURIFromPath(vCDQueryURIPath)
	queryURI := fmt.Sprintf("%s?%s%s", URI, templateQueryParams, templateName)

	if req, err = http.NewRequest("GET", queryURI, nil); err == nil {

		if err = s.AuthDecorate(req); err == nil {
			vappTemplate, err = s.queryAndParseResponse(req)
		}
	}
	return
}

func (s *VCDClient) getAbsoluteURIFromPath(uriPath string) (absoluteURI string) {
	absoluteURI = fmt.Sprintf("%s%s", s.BaseURI, uriPath)
	return
}

func (s *VCDClient) queryAndParseResponse(req *http.Request) (vappTemplate *VAppTemplateRecord, err error) {
	var (
		res  *http.Response
		body []byte
	)
	vappTemplate = new(VAppTemplateRecord)

	if res, err = s.client.Do(req); err == nil && res.StatusCode == QuerySuccessStatusCode {
		body, err = ioutil.ReadAll(res.Body)
		tmplt := QueryResultRecords{}
		xml.Unmarshal(body, &tmplt)
		*vappTemplate = firstElement(tmplt.VAppTemplateRecord)

	} else if res.StatusCode != QuerySuccessStatusCode && err == nil {
		err = ErrFailedQuery
	}
	return
}

func firstElement(va []VAppTemplateRecord) VAppTemplateRecord {
	return va[0]
}
