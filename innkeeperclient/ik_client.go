package innkeeperclient

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"strconv"

	"github.com/franela/goreq"
)

// New - create a new api client
func New(log logger) (clnt InnkeeperClient) {
	sPort := os.Getenv("INNKEEPER_PORT")
	port, err := strconv.Atoi(sPort)
	if err != nil {
		port = 5555
	}
	clnt = &IkClient{
		Host:     os.Getenv("INNKEEPER_HOST"),
		Port:     port,
		User:     os.Getenv("INNKEEPER_USER"),
		Password: os.Getenv("INNKEEPER_PASSWORD"),
		Log:      log,
	}
	return
}

// call -- generic call to the inkeeper endpoint
func (s *IkClient) call(path string, query interface{}, jsonResp interface{}) (err error) {
	res, err := goreq.Request{
		Uri:               "http://" + s.Host + ":" + strconv.Itoa(s.Port) + "/" + path,
		BasicAuthUsername: s.User,
		BasicAuthPassword: s.Password,
		QueryString:       query}.Do()
	if err == nil && res.StatusCode < 300 {
		res.Body.FromJsonTo(jsonResp)
		fmt.Println(res.Body.ToString())
	} else {
		s.Log.Println(res.Body.ToString())
		s.Log.Println(err)
		strerr, err := res.Body.ToString()
		if err == nil {
			err = errors.New(strerr)
		}
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
