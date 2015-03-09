package pezdispenser

import "github.com/go-martini/martini"

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

type LeaseTypeController struct {
	controllerBase
	version     string
	versionTree map[string]interface{}
}

func (s *LeaseTypeController) Delete() (post interface{}) {
	switch s.version {
	case ApiVersion1:
		post = s.postV1
	}
	return
}

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

type LeaseItemController struct {
	controllerBase
	version string
}

func (s *LeaseItemController) Delete() (post interface{}) {
	switch s.version {
	case ApiVersion1:
		post = s.postV1
	}
	return
}

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
