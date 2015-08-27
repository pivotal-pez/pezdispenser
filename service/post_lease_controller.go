package pezdispenser

import (
	"log"
	"net/http"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/pivotal-pez/pezdispenser/service/_integrations"
	"github.com/pivotal-pez/pezdispenser/skus"
)

//PostLeaseController - this is the controller to handle a get task call
func PostLeaseController(taskServiceURI string, collectionDialer integrations.CollectionDialer) martini.Handler {
	taskCollection := setupDB(collectionDialer, taskServiceURI, TaskCollectionName)
	availableSkus := map[string]skus.Sku{
		"2c.small": new(skus.Sku2CSmall),
	}
	return func(logger *log.Logger, r render.Render, req *http.Request) {
		lease := NewLease(taskCollection, availableSkus)
		statusCode, response := lease.Post(logger, req)
		r.JSON(statusCode, response)
	}
}

//DeleteLeaseController - this is the controller to handle a get task call
func DeleteLeaseController(taskServiceURI string, collectionDialer integrations.CollectionDialer) martini.Handler {
	taskCollection := setupDB(collectionDialer, taskServiceURI, TaskCollectionName)
	availableSkus := map[string]skus.Sku{
		"2c.small": new(skus.Sku2CSmall),
	}
	return func(logger *log.Logger, r render.Render, req *http.Request) {
		lease := NewLease(taskCollection, availableSkus)
		statusCode, response := lease.Delete(logger, req)
		r.JSON(statusCode, response)
	}
}
