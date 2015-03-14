package pezdispenser

import "github.com/go-martini/martini"

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
	}
	return
}

//LeaseTypeController - this is a controller for a lease for a specific type
type LeaseTypeController struct {
	controllerBase
	version string
}

//Post - this will return the versioned controller for a post rest call
func (s *LeaseTypeController) Post() (post interface{}) {
	switch s.version {
	case APIVersion1:
		post = s.postV1
	}
	return
}

func (s *LeaseTypeController) postV1(params martini.Params) (res string) {
	typeGUID := params[TypeGUID]
	res = typeGUID
	return
}

//LeaseItemController - this is a controller for a lease for a specific item
type LeaseItemController struct {
	controllerBase
	version string
}

//Delete - this will return the versioned controller for a delete rest call
func (s *LeaseItemController) Delete() (post interface{}) {
	switch s.version {
	case APIVersion1:
		post = s.deleteV1
	}
	return
}

//Post - this will return the versioned controller for a post rest call
func (s *LeaseItemController) Post() (post interface{}) {
	switch s.version {
	case APIVersion1:
		post = s.postV1
	}
	return
}

func (s *LeaseItemController) postV1(params martini.Params) (res string) {
	itemGUID := params[ItemGUID]
	res = itemGUID
	return
}

func (s *LeaseItemController) deleteV1(params martini.Params) (res string) {
	itemGUID := params[ItemGUID]
	res = itemGUID
	return
}
