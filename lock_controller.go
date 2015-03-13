package pezdispenser

import "github.com/go-martini/martini"

//NewLockController - returns a controller interface build using the version argument
func NewLockController(version string) (controller Controller) {
	controller = &LockController{
		version: version,
	}
	return
}

//LockController - a controller for locking leases when they hit their expiration date
type LockController struct {
	controllerBase
	version string
}

//Get - retuns a versioned controller for get requests
func (s *LockController) Get() (post interface{}) {
	switch s.version {
	case ApiVersion1:
		post = s.getV1
	}
	return
}

//Post - retuns a versioned controller for post requests
func (s *LockController) Post() (post interface{}) {
	switch s.version {
	case ApiVersion1:
		post = s.postV1
	}
	return
}

func (s *LockController) getV1(params martini.Params) string {
	return "Hello " + params[TypeGuid]
}

func (s *LockController) postV1(params martini.Params) string {
	return "Hello " + params[TypeGuid]
}
