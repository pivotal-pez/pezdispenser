package pezauth

import (
	"log"
	"net/http"

	"github.com/martini-contrib/render"
)

//NewValidateV1 - create a validation controller
func NewValidateV1(kg KeyGenerator) Controller {
	return &validateV1{
		keyGenerator: kg,
	}
}

func (s *validateV1) Get() interface{} {
	var handler ValidateGetHandler = func(log *log.Logger, r render.Render, req *http.Request) {
		responseBody := Response{}
		statusCode := SuccessStatus

		if key := req.Header.Get(HeaderKeyName); len(key) == GUIDLength {
			log.Println("checking key: ...-", key[:4])

			if _, val, err := s.keyGenerator.GetByKey(key); err == nil {
				log.Println("valid key match")
				responseBody.Payload = val
				responseBody.APIKey = key

			} else {
				log.Println(err)
				responseBody.ErrorMsg = err.Error()
				statusCode = FailureStatus
			}

		} else {
			log.Println(ErrInvalidKeyFormatMsg)
			responseBody.ErrorMsg = ErrInvalidKeyFormatMsg
			statusCode = FailureStatus
		}
		r.JSON(statusCode, responseBody)
	}
	return handler
}
