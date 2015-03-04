package main

import (
	"github.com/go-martini/martini"
	pez "github.com/pivotalservice/pezdispenser"
)

func main() {
	m := martini.Classic()
	pez.InitAuth(m)
	pez.InitRoutes(m)
	m.Run()
}
