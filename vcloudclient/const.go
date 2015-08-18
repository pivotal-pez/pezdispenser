package vcloudclient

import "errors"

const (

	//VCloudTokenHeaderName - response header name for the auth token
	VCloudTokenHeaderName = "X-Vcloud-Authorization"
	//AuthSuccessStatusCode - status code expected for a successful auth call to the vcd api
	AuthSuccessStatusCode = 200
	//VAppTemplateName - the name of the vapp template to seed our apps from by defualt
	VAppTemplateName = "vSphere6-base-pcfaas-0.9.2"
	//VCDAuthURIPath - path for the authentication rest calls
	VCDAuthURIPath = "/api/sessions"
	//VCDQueryURIPath - path for the query rest calls to vcd
	VCDQueryURIPath = "/api/query"
)

var (
	//ErrAuthFailure - error message returned for authentication responses not having a success statuscode
	ErrAuthFailure = errors.New("status code failure on authentication call to api")

	ErrNoTokenToApply = errors.New("no token to decorate the given request with")
)
