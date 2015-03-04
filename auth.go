package pezdispenser

import (
	"github.com/go-martini/martini"
)

type martiniUseable interface {
	Use(handler martini.Handler)
}

func InitAuth(m martiniUseable) {
	// do something here below copied from samples
	//m.Use(func(res http.ResponseWriter, req *http.Request) {
	//if req.Header.Get("X-API-KEY") != "secret123" {
	//res.WriteHeader(http.StatusUnauthorized)
	//}
	//})
}
