package pezdispenser

import (
	"encoding/json"

	"github.com/go-martini/martini"
)

//NewLeaseController - builds a new object of type controller from arguments
func NewLeaseController(version string, category int) (controller Controller) {
	switch category {
	case Item:
		controller = &LeaseItemController{
			version: version,
		}
	case Type:
		controller = &LeaseTypeController{
			version: version,
		}
	case List:
		controller = &LeaseListController{
			version: version,
		}
	}
	return
}

//LeaseTypeController - this is a controller for a lease for a specific type
type LeaseListController struct {
	controllerBase
	version string
}

//Get - this will return the versioned controller for a get rest call
func (s *LeaseListController) Get() (get interface{}) {
	get = s.getFnc
	return
}

func (s *LeaseListController) getFnc() (res string) {
	f := NewFinder()
	dispenserList := f.GetAll()
	resObj := ResponseMessage{Version: s.version}
	res = genericControlFormatter(resObj, func() (res []byte, err error) {
		r := []string{}

		for _, d := range dispenserList {
			r = append(r, d.GUID())
		}
		res, err = json.Marshal(r)
		return
	})
	return
}

//LeaseTypeController - this is a controller for a lease for a specific type
type LeaseTypeController struct {
	controllerBase
	version string
}

//Post - this will return the versioned controller for a post rest call
func (s *LeaseTypeController) Post() (post interface{}) {
	post = s.postFnc
	return
}

//Get - this will return the versioned controller for a get rest call
func (s *LeaseTypeController) Get() (get interface{}) {
	get = s.getFnc
	return
}

func (s *LeaseTypeController) getFnc(params martini.Params) (res string) {
	typeGUID := params[TypeGUID]
	f := NewFinder()
	dispenser := f.GetByTypeGUID(typeGUID)
	resObj := ResponseMessage{Version: s.version}
	res = genericControlFormatter(resObj, dispenser.Status)
	return
}

func (s *LeaseTypeController) postFnc(params martini.Params) (res string) {
	typeGUID := params[TypeGUID]
	f := NewFinder()
	dispenser := f.GetByTypeGUID(typeGUID)
	resObj := ResponseMessage{Version: s.version}
	res = genericControlFormatter(resObj, dispenser.Lease)
	return
}

//LeaseItemController - this is a controller for a lease for a specific item
type LeaseItemController struct {
	controllerBase
	version string
}

//Delete - this will return the versioned controller for a delete rest call
func (s *LeaseItemController) Delete() (post interface{}) {
	post = s.deleteFnc
	return
}

//Post - this will return the versioned controller for a post rest call
func (s *LeaseItemController) Post() (post interface{}) {
	post = s.postFnc
	return
}

//Get - this will return the versioned controller for a get rest call
func (s *LeaseItemController) Get() (get interface{}) {
	get = s.getFnc
	return
}

func (s *LeaseItemController) getFnc(params martini.Params) (res string) {
	inventoryGUID := params[ItemGUID]
	f := NewFinder()
	dispenser := f.GetByItemGUID(inventoryGUID)
	resObj := ResponseMessage{Version: s.version}
	res = genericControlFormatter(resObj, dispenser.Status)
	return
}

func (s *LeaseItemController) postFnc(params martini.Params) (res string) {
	inventoryGUID := params[ItemGUID]
	f := NewFinder()
	dispenser := f.GetByItemGUID(inventoryGUID)
	resObj := ResponseMessage{Version: s.version}
	res = genericControlFormatter(resObj, dispenser.Lease)
	return
}

func (s *LeaseItemController) deleteFnc(params martini.Params) (res string) {
	inventoryGUID := params[ItemGUID]
	f := NewFinder()
	dispenser := f.GetByItemGUID(inventoryGUID)
	resObj := ResponseMessage{Version: s.version}
	res = genericControlFormatter(resObj, dispenser.Unlease)
	return
}
