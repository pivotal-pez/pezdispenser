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

//Get - returns a versioned controller for get requests
func (s *LockController) Get() (post interface{}) {
	switch s.version {
	case APIVersion1:
		post = s.getV1
	}
	return
}

//Post - returns a versioned controller for post requests
func (s *LockController) Post() (post interface{}) {
	switch s.version {
	case APIVersion1:
		post = s.postV1
	}
	return
}

func (s *LockController) getV1(params martini.Params) (res string) {
	inventoryGUID := params[ItemGUID]
	res = inventoryGUID
	return
}

func (s *LockController) postV1(params martini.Params) (res string) {
	inventoryGUID := params[ItemGUID]
	res = inventoryGUID
	return
}
