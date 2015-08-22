package pezdispenser

import (
	"log"

	"labix.org/v2/mgo"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/pivotal-pez/pezdispenser/service/_integrations"
)

//GetTaskByIDController - this is the controller to handle a get task call
func GetTaskByIDController(taskServiceURI string, collectionDialer integrations.CollectionDialer) martini.Handler {
	return func(params martini.Params, log *log.Logger, r render.Render) {
		var (
			err        error
			dialInfo   *mgo.DialInfo
			response   interface{} = &err
			statusCode             = FailureStatusResponseTaskByID
			collection integrations.Collection
		)
		defer func() {
			if collection != nil {
				collection.Close()
			}
		}()
		task := new(Task)
		taskID := params["id"]

		if dialInfo, err = ParseURL(taskServiceURI); err == nil {

			if collection, err = collectionDialer(taskServiceURI, dialInfo.Database, TaskCollectionName); err == nil {
				err = collection.FindOne(taskID, task)
				statusCode = SuccessStatusResponseTaskByID
				response = task
			}
		}
		r.JSON(statusCode, response)
	}
}
