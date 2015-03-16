package pezdispenser

import (
	"encoding/json"

	"github.com/go-martini/martini"
)

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
	var (
		err  error
		resB []byte
	)
	inventoryGUID := params[ItemGUID]
	f := NewFinder()
	dispenser := f.GetByItemGUID(inventoryGUID)
	resObj := ResponseMessage{Version: s.version}

	if stat, err := dispenser.Status(); err == nil {
		resObj.Status = SuccessStatus
		resObj.Body = stat

	} else {
		resObj.Status = FailureStatus
	}

	if resB, err = json.Marshal(resObj); err != nil {
		res = err.Error()
	} else {
		res = string(resB[:])
	}
	return
}

func (s *LockController) postFnc(params martini.Params) (res string) {
	var (
		err  error
		resB []byte
	)
	inventoryGUID := params[ItemGUID]
	f := NewFinder()
	dispenser := f.GetByItemGUID(inventoryGUID)
	resObj := ResponseMessage{Version: s.version}

	if stat, err := dispenser.Lock(); err == nil {
		resObj.Status = SuccessStatus
		resObj.Body = stat

	} else {
		resObj.Status = FailureStatus
	}

	if resB, err = json.Marshal(resObj); err != nil {
		res = err.Error()
	} else {
		res = string(resB[:])
	}
	return
}
