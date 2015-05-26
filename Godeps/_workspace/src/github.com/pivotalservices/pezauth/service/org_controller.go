package pezauth

import (
	"errors"
	"log"

	"github.com/fatih/structs"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/oauth2"
	"github.com/martini-contrib/render"
)

const (
	//EmailFieldName - fieldname for email
	EmailFieldName = "email"
)

var (
	//ErrNoMatchInStore - error when there is no matching org in the datastore
	ErrNoMatchInStore = errors.New("Could not find a matching user org or connection failure")
	//ErrCanNotCreateOrg - error when we can not create an org
	ErrCanNotCreateOrg = errors.New("Could not create a new org")
	//ErrCanNotAddOrgRec - error when we can not add a new org record to the datastore
	ErrCanNotAddOrgRec = errors.New("Could not add a new org record")
	//ErrCantCallAcrossUsers - error when a user is trying to update a user record other than their own
	ErrCantCallAcrossUsers = errors.New("user calling another users endpoint")
)

type (
	//OrgGetHandler - func signature of org get handler
	OrgGetHandler func(params martini.Params, log *log.Logger, r render.Render, tokens oauth2.Tokens)
	//OrgPutHandler - func signature of org put handler
	OrgPutHandler func(params martini.Params, log *log.Logger, r render.Render, tokens oauth2.Tokens)
)

type (
	//Persistence - interface to a persistence store of some kind
	Persistence interface {
		FindOne(query interface{}, result interface{}) (err error)
		Upsert(selector interface{}, update interface{}) (err error)
	}
	//PivotOrg - struct for pivot org record
	PivotOrg struct {
		Email   string
		OrgName string
		OrgGUID string
	}
	orgController struct {
		Controller
		store      Persistence
		authClient AuthRequestCreator
	}
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
