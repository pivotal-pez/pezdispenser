package pezdispenser

import "github.com/go-martini/martini"

func GetInfoController() martini.Handler {
	return func() string {
		return "the dispenser service will give you candy"
	}
}
