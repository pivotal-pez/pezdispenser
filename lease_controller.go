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

//Delete - this will return the versioned controller for a delete rest call
func (s *LeaseTypeController) Delete() (post interface{}) {
	switch s.version {
	case ApiVersion1:
		post = s.postV1
	}
	return
}

//Post - this will return the versioned controller for a post rest call
func (s *LeaseTypeController) Post() (post interface{}) {
	switch s.version {
	case ApiVersion1:
		post = s.postV1
	}
	return
}

func (s *LeaseTypeController) postV1(params martini.Params) string {
	return "Hello " + params[TypeGuid]
}

func (s *LeaseTypeController) deleteV1(params martini.Params) string {
	return "Hello " + params[TypeGuid]
}

//LeaseItemController - this is a controller for a lease for a specific item
type LeaseItemController struct {
	controllerBase
	version string
}

//Delete - this will return the versioned controller for a delete rest call
func (s *LeaseItemController) Delete() (post interface{}) {
	switch s.version {
	case ApiVersion1:
		post = s.postV1
	}
	return
}

//Post - this will return the versioned controller for a post rest call
func (s *LeaseItemController) Post() (post interface{}) {
	switch s.version {
	case ApiVersion1:
		post = s.postV1
	}
	return
}

func (s *LeaseItemController) postV1(params martini.Params) string {
	return "Hello " + params[TypeGuid]
}

func (s *LeaseItemController) deleteV1(params martini.Params) string {
	return "Hello " + params[TypeGuid]
}
