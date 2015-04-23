package pezauth

import (
	"log"

	"github.com/martini-contrib/oauth2"
	"github.com/martini-contrib/render"
)

//MeGetHandler - a get control handler for me requests
type (
	MeGetHandler func(log *log.Logger, r render.Render, tokens oauth2.Tokens)
)

//NewMeController - a controller for me requests
func NewMeController() Controller {
	return new(meController)
}

type meController struct {
	Controller
}

//Get - get a get handler for authkeyv1
func (s *meController) Get() interface{} {
	var handler MeGetHandler = func(log *log.Logger, r render.Render, tokens oauth2.Tokens) {
		userInfo := GetUserInfo(tokens)
		log.Println("getting userInfo: ", userInfo)
		genericResponseFormatter(r, "", userInfo, nil)
	}
	return handler
}
