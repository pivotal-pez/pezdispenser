package vcloudclient

import "errors"

const (

	//VCloudTokenHeaderName - response header name for the auth token
	VCloudTokenHeaderName = "X-Vcloud-Authorization"
	//AuthSuccessStatusCode - status code expected for a successful auth call to the vcd api
	AuthSuccessStatusCode = 200
	//QuerySuccessStatusCode - a query call to vcd api was successful
	QuerySuccessStatusCode = 200
	vCDAuthURIPath         = "/api/sessions"
	vCDQueryURIPath        = "/api/query"
	templateQueryParams    = "type=vAppTemplate&filter=name=="
)

var (
	//ErrAuthFailure - error message returned for authentication responses not having a success statuscode
	ErrAuthFailure = errors.New("status code failure on authentication call to api")
	//ErrNoTokenToApply - error when there is no auth token
	ErrNoTokenToApply = errors.New("no token to decorate the given request with")
	//ErrFailedQuery - query to vcd api failed returning non 200 status code
	ErrFailedQuery = errors.New("invalid response code from query api call")
)
