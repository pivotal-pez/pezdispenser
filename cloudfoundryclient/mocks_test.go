package cloudfoundryclient_test

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/xchapter7x/cloudcontroller-client"
)

type mockRequestDecorator struct {
	doer ccclient.ClientDoer
}

func (s *mockRequestDecorator) CreateAuthRequest(verb, requestURL, path string, args interface{}) (*http.Request, error) {
	return nil, nil
}

func (s *mockRequestDecorator) CCTarget() string {
	return ""
}

func (s *mockRequestDecorator) HttpClient() ccclient.ClientDoer {
	return s.doer
}

func (s *mockRequestDecorator) Login() (*ccclient.Client, error) {
	return nil, nil
}

type mockClientDoer struct {
	res *http.Response
	err error
}

func (s *mockClientDoer) Do(req *http.Request) (*http.Response, error) {
	return s.res, s.err
}

type mockLog struct{}

func (s *mockLog) Println(i ...interface{}) {
	fmt.Println(i...)
}

type nopCloser struct {
	io.Reader
}

func (nopCloser) Close() error { return nil }

func mockHttpResponse(body string, statusCode int) *http.Response {
	return &http.Response{
		StatusCode: statusCode,
		Body:       nopCloser{bytes.NewBufferString(body)},
	}
}

var (
	mockSuccessInfoStatusCode   = 200
	mockSuccessInfoResponseBody = `{"name":"vcap","build":"2222","support":"http://support.cloudfoundry.com","version":2,"description":"Cloud Foundry sponsored by Pivotal","authorization_endpoint":"http://localhost:8080/uaa","token_endpoint":"http://localhost:8080/uaa","min_cli_version":null,"min_recommended_cli_version":null,"api_version":"2.27.0","logging_endpoint":"ws://loggregator.vcap.me:80"}`
	mockSuccessUserStatusCode   = 200
	mockSuccessUserResponseBody = `{"resources":[{"id":"123456","userName":"testuser"}],"startIndex":1,"itemsPerPage":100,"totalResults":1,"schemas":["urn:scim:schemas:core:1.0"]}`
	mockSuccessRoleStatusCode   = 201
	mockSuccessRoleResponseBody = `{"metadata": {"guid": "13b9e0fb-0ae7-4af0-bd07-7739d05caad3","url": "/v2/organizations/13b9e0fb-0ae7-4af0-bd07-7739d05caad3","created_at": "2015-05-15T16:53:36Z","updated_at": null},"entity": {"name": "name-843","billing_enabled": false,"quota_definition_guid": "deab75fd-5a6e-46d0-89da-349d57ed9e09","status": "active","quota_definition_url": "/v2/quota_definitions/deab75fd-5a6e-46d0-89da-349d57ed9e09","spaces_url": "/v2/organizations/13b9e0fb-0ae7-4af0-bd07-7739d05caad3/spaces","domains_url": "/v2/organizations/13b9e0fb-0ae7-4af0-bd07-7739d05caad3/domains","private_domains_url": "/v2/organizations/13b9e0fb-0ae7-4af0-bd07-7739d05caad3/private_domains","users_url": "/v2/organizations/13b9e0fb-0ae7-4af0-bd07-7739d05caad3/users","managers_url": "/v2/organizations/13b9e0fb-0ae7-4af0-bd07-7739d05caad3/managers","billing_managers_url": "/v2/organizations/13b9e0fb-0ae7-4af0-bd07-7739d05caad3/billing_managers","auditors_url": "/v2/organizations/13b9e0fb-0ae7-4af0-bd07-7739d05caad3/auditors","app_events_url": "/v2/organizations/13b9e0fb-0ae7-4af0-bd07-7739d05caad3/app_events","space_quota_definitions_url": "/v2/organizations/13b9e0fb-0ae7-4af0-bd07-7739d05caad3/space_quota_definitions"}}`
	mockSuccessOrgStatusCode    = 201
	mockSuccessOrgResponseBody  = `{"metadata": {"guid": "1e2bae2c-459e-4ad8-b1cb-ffc09d209b32","url": "/v2/organizations/1e2bae2c-459e-4ad8-b1cb-ffc09d209b32","created_at": "2015-05-15T16:53:35Z","updated_at": null},"entity": {"name": "my-org-name","billing_enabled": false,"quota_definition_guid": "b97156be-2d35-43ba-a358-9e1b04d6a877","status": "active","quota_definition_url": "/v2/quota_definitions/b97156be-2d35-43ba-a358-9e1b04d6a877","spaces_url": "/v2/organizations/1e2bae2c-459e-4ad8-b1cb-ffc09d209b32/spaces","domains_url": "/v2/organizations/1e2bae2c-459e-4ad8-b1cb-ffc09d209b32/domains","private_domains_url": "/v2/organizations/1e2bae2c-459e-4ad8-b1cb-ffc09d209b32/private_domains","users_url": "/v2/organizations/1e2bae2c-459e-4ad8-b1cb-ffc09d209b32/users","managers_url": "/v2/organizations/1e2bae2c-459e-4ad8-b1cb-ffc09d209b32/managers","billing_managers_url": "/v2/organizations/1e2bae2c-459e-4ad8-b1cb-ffc09d209b32/billing_managers","auditors_url": "/v2/organizations/1e2bae2c-459e-4ad8-b1cb-ffc09d209b32/auditors","app_events_url": "/v2/organizations/1e2bae2c-459e-4ad8-b1cb-ffc09d209b32/app_events","space_quota_definitions_url": "/v2/organizations/1e2bae2c-459e-4ad8-b1cb-ffc09d209b32/space_quota_definitions"}}`
)
