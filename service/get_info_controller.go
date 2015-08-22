package pezdispenser

import "github.com/go-martini/martini"

//GetInfoController - this is the controller to handle a info call to the api
func GetInfoController() martini.Handler {
	return func() string {
		return "the dispenser service will give you candy"
	}
}
