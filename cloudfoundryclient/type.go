package cloudfoundryclient

import (
	"net/http"

	"github.com/xchapter7x/cloudcontroller-client"
)

type (
	//RestRunner - runs a rest call
	RestRunner struct {
		Verb              string
		URL               string
		Data              interface{}
		Path              string
		SuccessStatusCode int
		OnSuccess         func(*http.Response)
		OnFailure         func(*http.Response, error)
		RequestDecorator  AuthRequestCreator
		Logger            logger
	}
	//AuthRequestCreator - creates auth decorated http request objects
	AuthRequestCreator interface {
		CreateAuthRequest(verb, requestURL, path string, args interface{}) (*http.Request, error)
		CCTarget() string
		HttpClient() ccclient.ClientDoer
		Login() (*ccclient.Client, error)
	}
	//CloudFoundryClient - interface for a cloud foundry client
	CloudFoundryClient interface {
		QueryAPIInfo() (*CloudFoundryAPIInfo, error)
		QueryUserGUID(username string) (string, error)
		AddRole(rolePathPrefix string, targetGUID string, roleType string, userGUID string) error
		AddOrg(orgName string) (orgGUID string, err error)
		AddSpace(spaceName string, orgGUID string) (spaceGUID string, err error)
		AddUser(username string) error
		RemoveOrg(orgGUID string) (err error)
		QueryUsers(int, int, string, string) (userList map[string]interface{}, err error)
	}
	//CFClient - cloud foundry api client struct
	CFClient struct {
		RequestDecorator AuthRequestCreator
		Info             *CloudFoundryAPIInfo
		Log              logger
	}
	logger interface {
		Println(...interface{})
	}
	//CloudFoundryAPIInfo - info response object from cc info endpoint
	CloudFoundryAPIInfo struct {
		Name                     string `json:"name"`
		Build                    string `json:"build"`
		Support                  string `json:"support"`
		Version                  int    `json:"version"`
		Description              string `json:"description"`
		AuthorizationEndpoint    string `json:"authorization_endpoint"`
		TokenEndpoint            string `json:"token_endpoint"`
		MinCLIVersion            string `json:"min_cli_version"`
		MinRecommendedCLIVersion string `json:"min_recommended_cli_version"`
		APIVersion               string `json:"api_version"`
		LoggingEndpoint          string `json:"logging_endpoint"`
	}
	//APIResponse - cc http response object
	APIResponse struct {
		Metadata APIMetadata            `json:"metadata"`
		Entity   map[string]interface{} `json:"entity"`
	}
	//APIMetadata = cc http response metadata
	APIMetadata struct {
		GUID      string `json:"guid"`
		URL       string `json:"url"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}
	//APIResponseList - a list of resources or apiresponse objects
	APIResponseList struct {
		StartIndex   int           `json:"startIndex"`
		ItemsPerPage int           `json:"itemsPerPage"`
		TotalResults int           `json:"totalResults"`
		Resources    []APIResponse `json:"resources"`
	}
	//UserAPIResponse - the user api response object
	UserAPIResponse struct {
		Resources []map[string]interface{}
	}
)
