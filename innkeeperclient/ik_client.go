package innkeeperclient

import (
	"errors"
	"net/http"
	"net/url"

	"github.com/franela/goreq"
)

// New - create a new api client
func New(uri string, user string, password string) InnkeeperClient {
	return &IkClient{
		URI:      uri,
		User:     user,
		Password: password,
	}
}

func isStatusOK(statusCode int) bool {
	return (statusCode == http.StatusOK || statusCode == http.StatusCreated || statusCode == http.StatusAccepted)
}

// call -- generic call to the inkeeper endpoint
func (s *IkClient) Call(path string, query interface{}, jsonResp interface{}) (err error) {
	res, err := goreq.Request{
		Insecure:          true,
		Uri:               s.URI + "/" + path,
		BasicAuthUsername: s.User,
		BasicAuthPassword: s.Password,
		QueryString:       query}.Do()
	if err == nil {

		if isStatusOK(res.StatusCode) {
			err = res.Body.FromJsonTo(jsonResp)

		} else {
			strerr, _ := res.Body.ToString()
			err = errors.New(strerr)
		}
	}
	return
}

//DeProvisionHost - make a deprovision call to innkeeper for a given requestID
func (s *IkClient) DeProvisionHost(requestID string) (err error) {
	resp := new(GetStatusResponse)
	qp := url.Values{}
	qp.Add("requestid", requestID)
	err = s.Call(RouteDeProvisionHost, qp, resp)
	return
}

//GetStatus --
func (s *IkClient) GetStatus(requestID string) (resp *GetStatusResponse, err error) {
	resp = new(GetStatusResponse)
	qp := url.Values{}
	qp.Add("requestid", requestID)
	err = s.Call(RouteGetStatus, qp, resp)
	return
}

// GetTenants -- /api/v1/GetTenants get current tenants
func (s *IkClient) GetTenants() (info GetTenantsResponse, err error) {
	return
}

// ProvisionHost -- given info provision a host in inkeeper
// "http://pez-app.core.pao.pez.pivotal.io:5555/api/v1/ProvisionHost?geo_loc=PAO&sku=4D.lowmem.R7&os=esxi60u2&count=1&feature=&tenantid=pez-stage"
func (s *IkClient) ProvisionHost(sku string, tenantid string) (info *ProvisionHostResponse, err error) {
	info = new(ProvisionHostResponse)
	qp := url.Values{}
	qp.Add("sku", sku)
	qp.Add("tenantid", tenantid)
	err = s.Call(RouteProvisionHost, qp, info)
	return
}
