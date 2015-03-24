package pezdispenser

import (
	"encoding/json"
	"errors"
)

//Different contexts for rest calls
const (
	Type = iota
	Item
	List
	SuccessStatus = "success"
	FailureStatus = "error"
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

//ResponseMessage - this is the structure of a response from any call to a controller
type ResponseMessage struct {
	Version string
	Body    []byte
	Status  string
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

func genericControlFormatter(resObj ResponseMessage, action func() ([]byte, error)) (res string) {
	var (
		err  error
		resB []byte
	)

	if stat, err := action(); err == nil {
		resObj.Status = SuccessStatus
		resObj.Body = stat

	} else {
		resObj.Status = FailureStatus
	}

	if resB, err = json.Marshal(resObj); err != nil {
		res = err.Error()
	} else {
		res = string(resB[:])
	}
	return
}
