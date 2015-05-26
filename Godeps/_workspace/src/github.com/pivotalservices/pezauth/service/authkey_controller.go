package pezauth

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/oauth2"
	"github.com/martini-contrib/render"
)

//Authentication Handler vars
var (
	ErrInvalidCallerEmail = errors.New("Invalid user token for your requested action")
)

type (
	//AuthPutHandler - auth control handler for put calls
	AuthPutHandler func(params martini.Params, log *log.Logger, r render.Render, tokens oauth2.Tokens)
	//AuthPostHandler - auth control handler for post calls
	AuthPostHandler func(params martini.Params, log *log.Logger, r render.Render, tokens oauth2.Tokens)
	//AuthGetHandler - auth control handler for get calls
	AuthGetHandler func(params martini.Params, log *log.Logger, r render.Render, tokens oauth2.Tokens)
	//AuthDeleteHandler - auth control handler for delete calls
	AuthDeleteHandler func(params martini.Params, log *log.Logger, r render.Render, tokens oauth2.Tokens)
)

//NewAuthKeyV1 - get an instance of a V1 authkey controller
func NewAuthKeyV1(kg KeyGenerator) Controller {
	return &authKeyV1{
		keyGen: kg,
	}
}

type authKeyV1 struct {
	Controller
	keyGen KeyGenerator
}

//Put - get a put handler for authkeyv1
func (s *authKeyV1) Put() interface{} {
	var handler AuthPutHandler = func(params martini.Params, log *log.Logger, r render.Render, tokens oauth2.Tokens) {
		var (
			err    error
			apikey string
		)
		log.Println("executing the put handler")
		username := params[UserParam]
		userInfo := GetUserInfo(tokens)
		details, _ := json.Marshal(userInfo)

		NewUserMatch().
			UserInfo(userInfo).
			UserName(username).
			OnSuccess(func() {
			log.Println("getting userInfo: ", userInfo)

			if err = s.keyGen.Delete(username); err != nil {
				log.Println("keyGen.Delete error: ", err)
			}

			if err = s.keyGen.Create(username, string(details[:])); err != nil {
				log.Println("keyGen.Create error: ", err)
			}

			if apikey, err = s.keyGen.Get(username); err != nil {
				log.Println("keyGen.Get error: ", err)
			}
		}).
			OnFailure(func() {
			err = ErrInvalidCallerEmail
			log.Println("invalid user token error: ", err)
		}).Run()

		genericResponseFormatter(r, apikey, userInfo, err)
	}
	return handler
}

//Get - get a get handler for authkeyv1
func (s *authKeyV1) Get() interface{} {
	var handler AuthGetHandler = func(params martini.Params, log *log.Logger, r render.Render, tokens oauth2.Tokens) {
		var (
			apikey string
			err    error
		)
		username := params[UserParam]
		userInfo := GetUserInfo(tokens)
		log.Println("getting userInfo: ", userInfo)

		NewUserMatch().
			UserInfo(userInfo).
			UserName(username).
			OnSuccess(func() {

			if apikey, err = s.keyGen.Get(username); err != nil {
				log.Println("keyGen.Get error:", err)
			}
		}).
			OnFailure(func() {
			err = ErrInvalidCallerEmail
			log.Println("invalid user token error: ", err)
		}).Run()

		genericResponseFormatter(r, apikey, userInfo, err)
	}
	return handler
}

//Delete - get a delete handler for authkeyv1
func (s *authKeyV1) Delete() interface{} {
	var handler AuthDeleteHandler = func(params martini.Params, log *log.Logger, r render.Render, tokens oauth2.Tokens) {
		var err error
		username := params[UserParam]
		log.Println("deleting apikey for: ", username)
		userInfo := GetUserInfo(tokens)

		NewUserMatch().
			UserInfo(userInfo).
			UserName(username).
			OnSuccess(func() {

			if err = s.keyGen.Delete(username); err == nil {
				log.Println("key deleted for: ", username)

			} else {
				log.Println("key delete failed: ", username, err.Error())
			}
		}).
			OnFailure(func() {
			err = ErrInvalidCallerEmail
			log.Println("invalid user token error: ", err)
		}).Run()

		genericResponseFormatter(r, "", userInfo, err)
	}
	return handler
}
