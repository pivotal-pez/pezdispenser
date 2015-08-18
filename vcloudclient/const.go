package vcloudclient

import "errors"

const (

	//VCloudTokenHeaderName - response header name for the auth token
	VCloudTokenHeaderName = "X-Vcloud-Authorization"
	//AuthSuccessStatusCode - status code expected for a successful auth call to the vcd api
	AuthSuccessStatusCode = 200
	//QuerySuccessStatusCode - a query call to vcd api was successful
	QuerySuccessStatusCode = 200
	//DeployVappSuccessStatusCode - successful status code for vapp deploy
	DeployVappSuccessStatusCode = 201
	vCDAuthURIPath              = "/api/sessions"
	vCDQueryURIPath             = "/api/query"
	vCDVAppDeploymentPath       = "/action/instantiateVAppTemplate"
	templateQueryParams         = "type=vAppTemplate&filter=name=="
	vAppDeploymentContentType   = "application/vnd.vmware.vcloud.instantiateVAppTemplateParams+xml; charset=ISO-8859-1"
	vAppDeploymentPayload       = `
		<InstantiateVAppTemplateParams 
		xmlns="http://www.vmware.com/vcloud/v1.5"
		xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
		xmlns:ovf="http://schemas.dmtf.org/ovf/envelope/1"
		name="%s"
		deploy="true"
		powerOn="true">
			<Description>PEZ PCFaaS</Description>
			<InstantiationParams></InstantiationParams>
			<Source 
			href="%s" />
		</InstantiateVAppTemplateParams>
		`
)

var (
	//ErrAuthFailure - error message returned for authentication responses not having a success statuscode
	ErrAuthFailure = errors.New("status code failure on authentication call to api")
	//ErrNoTokenToApply - error when there is no auth token
	ErrNoTokenToApply = errors.New("no token to decorate the given request with")
	//ErrFailedQuery - query to vcd api failed returning non 200 status code
	ErrFailedQuery = errors.New("invalid response code from query api call")
)
