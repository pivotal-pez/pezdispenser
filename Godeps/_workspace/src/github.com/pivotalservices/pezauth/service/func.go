package pezauth

import (
	"encoding/json"

	"github.com/martini-contrib/render"
)

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
