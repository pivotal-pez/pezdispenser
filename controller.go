package pezdispenser

import "errors"

const (
	Type = iota
	Item
)

var (
	UndefinedPostError   = errors.New("You must define a Post() function for your struct that extends Controller")
	UndefinedGetError    = errors.New("You must define a Get() function for your struct that extends Controller")
	UndefinedDeleteError = errors.New("You must define a Delete() function for your struct that extends Controller")
)

type Controller interface {
	Get() interface{}
	Post() interface{}
	Delete() interface{}
}

type controllerBase struct {
}

func (s *controllerBase) Post() (i interface{}) {
	panic(UndefinedPostError)
	return
}

func (s *controllerBase) Delete() (i interface{}) {
	panic(UndefinedDeleteError)
	return
}

func (s *controllerBase) Get() (i interface{}) {
	panic(UndefinedGetError)
	return
}
