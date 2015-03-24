package main

import (
	"github.com/go-martini/martini"
	pez "github.com/pivotalservices/pezdispenser/service"
)

func main() {
	m := martini.Classic()
	pez.InitAuth(m)
	pez.InitRoutes(m)
	m.Run()
}
