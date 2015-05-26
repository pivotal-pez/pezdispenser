package pezauth

import (
	"encoding/json"
	"net/http"

	"github.com/martini-contrib/render"
	"github.com/xchapter7x/cloudcontroller-client"
)

const (
	//FailureStatus - failure response status from our unauthenticated rest endpoints
	FailureStatus = 403
	//SuccessStatus - success response status from our authenticated rest endpoints
	SuccessStatus = 200
)

//Controller - interface of a base controller
type Controller interface {
	Put() interface{}
	Post() interface{}
	Get() interface{}
	Delete() interface{}
}

//AuthRequestCreator - interface to an object which can decorate a request with auth tokens
type AuthRequestCreator interface {
	CreateAuthRequest(verb, requestURL, path string, args interface{}) (*http.Request, error)
	CCTarget() string
	HttpClient() ccclient.ClientDoer
	Login() (*ccclient.Client, error)
}

func genericResponseFormatter(r render.Render, apikey string, payload map[string]interface{}, extErr error) {
	var (
		statusCode int
		err        error
		res        Response
	)

	if extErr != nil {
		statusCode = FailureStatus
		res = Response{
			ErrorMsg: extErr.Error(),
		}

	} else {

		if _, err = json.Marshal(payload); err != nil {
			statusCode = FailureStatus
			res = Response{
				ErrorMsg: err.Error(),
			}

		} else {
			statusCode = SuccessStatus
			res = Response{
				APIKey:  apikey,
				Payload: payload,
			}
		}
	}
	r.JSON(statusCode, res)
}
