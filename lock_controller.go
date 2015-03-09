package pezdispenser

import "github.com/go-martini/martini"

func NewLockController(version string) (controller Controller) {
	controller = &LockController{
		version: version,
	}
	return
}

type LockController struct {
	controllerBase
	version string
}

func (s *LockController) Get() (post interface{}) {
	switch s.version {
	case ApiVersion1:
		post = s.getV1
	}
	return
}

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
