package pezdispenser

import (
	"log"
	"net/http"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/pivotal-pez/pezdispenser/service/integrations"
	"github.com/xchapter7x/lo"
)

//PostLeaseController - this is the controller to handle a get task call
func PostLeaseController() martini.Handler {
	return func(logger *log.Logger, r render.Render, req *http.Request, taskCollection integrations.Collection) {
		lease := NewLease(taskCollection, GetAvailableInventory(taskCollection))
		statusCode, response := lease.Post(logger, req)
		lo.G.Debug("statuscode: ", statusCode)
		lo.G.Debug("response: ", response)
		r.JSON(statusCode, response)
	}
}

//DeleteLeaseController - this is the controller to handle a get task call
func DeleteLeaseController() martini.Handler {
	return func(logger *log.Logger, r render.Render, req *http.Request, taskCollection integrations.Collection) {
		lease := NewLease(taskCollection, GetAvailableInventory(taskCollection))
		statusCode, response := lease.Delete(logger, req)
		lo.G.Debug("statuscode: ", statusCode)
		lo.G.Debug("response: ", response)
		r.JSON(statusCode, response)
	}
}
