package pezdispenser

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/auth"
)

const (
	user = "admin"
	pass = "temporary"
)

type martiniUseable interface {
	Use(handler martini.Handler)
}

//InitAuth - initializes authentication middleware for controllers
func InitAuth(m martiniUseable) {
	m.Use(auth.BasicFunc(func(username, password string) bool {
		return auth.SecureCompare(username, user) && auth.SecureCompare(password, pass)
	}))
}
