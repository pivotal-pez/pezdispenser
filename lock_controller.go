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
	post = s.getFnc
	return
}

//Post - returns a versioned controller for post requests
func (s *LockController) Post() (post interface{}) {
	post = s.postFnc
	return
}

func (s *LockController) getFnc(params martini.Params) (res string) {
	inventoryGUID := params[ItemGUID]
	f := NewFinder()
	dispenser := f.GetByItemGUID(inventoryGUID)
	resObj := ResponseMessage{Version: s.version}
	res = genericControlFormatter(resObj, dispenser.Status)
	return
}

func (s *LockController) postFnc(params martini.Params) (res string) {
	inventoryGUID := params[ItemGUID]
	f := NewFinder()
	dispenser := f.GetByItemGUID(inventoryGUID)
	resObj := ResponseMessage{Version: s.version}
	res = genericControlFormatter(resObj, dispenser.Lock)
	return
}
