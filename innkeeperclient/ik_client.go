package innkeeperclient

import (
	"errors"
	"net/url"

	"github.com/franela/goreq"
	"github.com/xchapter7x/lo"
)

// New - create a new api client
func New(uri string, user string, password string) InnkeeperClient {
	return &IkClient{
		URI:      uri,
		User:     user,
		Password: password,
	}
}

// call -- generic call to the inkeeper endpoint
func (s *IkClient) call(path string, query interface{}, jsonResp interface{}) (err error) {
	res, err := goreq.Request{
		Uri:               s.URI + "/" + path,
		BasicAuthUsername: s.User,
		BasicAuthPassword: s.Password,
		QueryString:       query}.Do()

	if err != nil {
		lo.G.Error(err.Error())
		return err
	}

	if res.StatusCode < 300 {
		err = res.Body.FromJsonTo(jsonResp)
		if err != nil {
			lo.G.Error(err.Error())
		}
	} else {
		lo.G.Debug(res.Body.ToString())
		strerr, err := res.Body.ToString()
		if err == nil {
			err = errors.New(strerr)
		}
		lo.G.Error(err.Error())
	}
	return
}

//GetStatus --
func (s *IkClient) GetStatus(requestID string) (resp *GetStatusResponse, err error) {
	return
}

// GetTenants -- /api/v1/GetTenants get current tenants
func (s *IkClient) GetTenants() (info GetTenantsResponse, err error) {
	err = s.call("api/v1/GetTenants", nil, &info)
	return
}

// ProvisionHost -- given info provision a host in inkeeper
// "http://pez-app.core.pao.pez.pivotal.io:5555/api/v1/ProvisionHost?geo_loc=PAO&sku=4D.lowmem.R7&os=esxi60u2&count=1&feature=&tenantid=pez-stage"
func (s *IkClient) ProvisionHost(sku string, tenantid string) (info *ProvisionHostResponse, err error) {
	info = new(ProvisionHostResponse)
	qp := url.Values{}
	qp.Add("sku", sku)
	qp.Add("tenantid", tenantid)
	err = s.call("api/v1/ProvisionHost", qp, info)
	return
}
