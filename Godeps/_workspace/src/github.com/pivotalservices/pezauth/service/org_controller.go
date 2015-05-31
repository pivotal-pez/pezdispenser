package pezauth

import (
	"log"

	"github.com/fatih/structs"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/oauth2"
	"github.com/martini-contrib/render"
)

//NewOrgController - a controller for me requests
func NewOrgController(c Persistence, authClient AuthRequestCreator) Controller {
	return &orgController{
		store:      c,
		authClient: authClient,
	}
}

//Get - get a get handler for org management
func (s *orgController) Get() interface{} {
	var handler OrgGetHandler = func(params martini.Params, log *log.Logger, r render.Render, tokens oauth2.Tokens) {
		username := params[UserParam]
		org := NewOrg(username, log, tokens, s.store, s.authClient)
		result, err := org.Show()
		genericResponseFormatter(r, "", structs.Map(result), err)
	}
	return handler
}

//Put - get a get handler for org management
func (s *orgController) Put() interface{} {
	var handler OrgPutHandler = func(params martini.Params, log *log.Logger, r render.Render, tokens oauth2.Tokens) {
		var (
			err             error
			payload         *PivotOrg
			responsePayload map[string]interface{}
		)
		username := params[UserParam]
		org := NewOrg(username, log, tokens, s.store, s.authClient)

		if _, err = org.Show(); err == ErrNoMatchInStore {
			payload, err = org.SafeCreate()
			responsePayload = structs.Map(payload)

		} else {
			err = ErrCanNotCreateOrg
		}
		genericResponseFormatter(r, "", responsePayload, err)
	}
	return handler
}
