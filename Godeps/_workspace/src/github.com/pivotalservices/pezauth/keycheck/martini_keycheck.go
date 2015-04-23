package keycheck

import (
	"log"
	"net/http"

	"github.com/go-martini/martini"
	"github.com/pivotalservices/pezauth/service"
)

const (
	//AuthFailStatus - failure status code
	AuthFailStatus = 403
	//AuthSuccessStatus - success status code
	AuthSuccessStatus = 200
	//KeyLength - valid length of a api-key
	KeyLength = pezauth.GUIDLength
)

var (
	//AuthFailureResponse - failure response for invalid key
	AuthFailureResponse = []byte(`{"error": "no valid key found"}`)
)

type (
	//APIKeyCheckHandler - type of our handler function
	APIKeyCheckHandler func(log *log.Logger, res http.ResponseWriter, req *http.Request)
)

//NewAPIKeyCheckMiddleware - creates a new instance of our middleware
func NewAPIKeyCheckMiddleware(url string) *APIKeyCheckMiddleware {
	keycheck := New(url)
	middleware := &APIKeyCheckMiddleware{Keycheck: keycheck}
	return middleware
}

//APIKeyCheckMiddleware - our middleware struct
type APIKeyCheckMiddleware struct {
	Keycheck KeyChecker
}

func badCheckCall(err error, res *http.Response) bool {
	return (err != nil || res.StatusCode != AuthSuccessStatus)
}

//Handler - returns the handler function as a martini.Handler type
func (s *APIKeyCheckMiddleware) Handler() martini.Handler {
	var handler APIKeyCheckHandler = func(log *log.Logger, res http.ResponseWriter, req *http.Request) {

		if key := req.Header.Get(HeaderKeyName); len(key) == KeyLength {

			if kcResponse, err := s.Keycheck.Check(key); badCheckCall(err, kcResponse) {
				log.Println("KeyAuth Failed: ", kcResponse, err)
				res.WriteHeader(AuthFailStatus)
				res.Write(AuthFailureResponse)

			} else {
				log.Println("KeyAuth Success: ", kcResponse)
			}

		} else {
			log.Println(AuthFailureResponse)
			res.WriteHeader(AuthFailStatus)
			res.Write(AuthFailureResponse)
		}
	}
	return handler
}
