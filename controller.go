package pezdispenser

import "errors"

//Different contexts for rest calls
const (
	Type = iota
	Item
)

//Definition of errors for controller
var (
	ErrUndefinedPost   = errors.New("You must define a Post() function for your struct that extends Controller")
	ErrUndefinedGet    = errors.New("You must define a Get() function for your struct that extends Controller")
	ErrUndefinedDelete = errors.New("You must define a Delete() function for your struct that extends Controller")
)

//Controller - This is a controller's interface
type Controller interface {
	Get() interface{}
	Post() interface{}
	Delete() interface{}
}

type controllerBase struct {
}

func (s *controllerBase) Post() (i interface{}) {
	panic(ErrUndefinedPost)
	return
}

func (s *controllerBase) Delete() (i interface{}) {
	panic(ErrUndefinedDelete)
	return
}

func (s *controllerBase) Get() (i interface{}) {
	panic(ErrUndefinedGet)
	return
}
