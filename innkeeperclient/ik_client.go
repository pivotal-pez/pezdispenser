package innkeeperclient

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"github.com/cloudfoundry-community/go-cfenv"
	"github.com/franela/goreq"
	"github.com/xchapter7x/lo"
)

// New - create a new api client
func New(log logger) (clnt InnkeeperClient) {
	if appEnv, err := cfenv.Current(); err == nil {

		if taskService, err := appEnv.Services.WithName("innkeeper-service"); err == nil {
				clnt = &IkClient{
			Uri: taskService.Credentials["uri"].(string),
			Password: taskService.Credentials["password"].(string),
			User: taskService.Credentials["user"].(string),
			Log: log,
			}

		} else {
			lo.G.Error("Experienced an error trying to grab innkeeper service binding information:", err.Error())
		}

	} else {
		lo.G.Error("error parsing current cfenv: ", err.Error())
	}
	return
}

// call -- generic call to the inkeeper endpoint
func (s *IkClient) call(path string, query interface{}, jsonResp interface{}) (err error) {
	res, err := goreq.Request{
		Uri:               s.Uri + "/" + path,
		BasicAuthUsername: s.User,
		BasicAuthPassword: s.Password,
		QueryString:       query}.Do()
	
	if err != nil{
		s.Log.Println(err)
		fmt.Println(err.Error())
		return err
	}
	
	if res.StatusCode < 300 {
		res.Body.FromJsonTo(jsonResp)
	} else {
		s.Log.Println(res.Body.ToString())
		strerr, err := res.Body.ToString()
		if err == nil {
			err = errors.New(strerr)
		}		
		s.Log.Println(err)
	}
	return
}

// GetTenants -- /api/v1/GetTenants get current tenants
func (s *IkClient) GetTenants() (info GetTenantsResponse, err error) {
	err = s.call("api/v1/GetTenants", nil, &info)
	return
}

// ProvisionHost -- given info provision a host in inkeeper
// "http://pez-app.core.pao.pez.pivotal.io:5555/api/v1/ProvisionHost?geo_loc=PAO&sku=4D.lowmem.R7&os=esxi60u2&count=1&feature=&tenantid=pez-stage"
func (s *IkClient) ProvisionHost(geoLoc string, sku string, count int, tenantid string, osarg string) (info ProvisionHostResponse, err error) {
	qp := url.Values{}
	qp.Add("goe_loc", geoLoc)
	qp.Add("sku", sku)
	qp.Add("count", strconv.Itoa(count))
	qp.Add("tenantid", tenantid)
	qp.Add("os", osarg)
	err = s.call("api/v1/ProvisionHost", qp, &info)
	return
}
