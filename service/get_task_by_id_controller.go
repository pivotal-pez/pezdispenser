package pezdispenser

import (
	"fmt"
	"log"

	"labix.org/v2/mgo"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/pivotal-pez/pezdispenser/service/_integrations"
)

//GetTaskByIDController - this is the controller to handle a get task call
func GetTaskByIDController(taskServiceURI string, collectionDialer integrations.CollectionDialer) martini.Handler {

	var (
		err      error
		dialInfo *mgo.DialInfo
	)
	if dialInfo, err = ParseURL(taskServiceURI); err == nil {
		log.Println("parsed uri successfully: ", taskServiceURI, dialInfo)

	} else {
		panic(fmt.Sprintf("can not parse given URI %s due to error: %s", taskServiceURI, err.Error()))
	}

	return func(params martini.Params, log *log.Logger, r render.Render) {
		var (
			err        error
			response   interface{} = &err
			statusCode             = FailureStatusResponseTaskByID
			collection integrations.Collection
			task       = new(Task)
			taskID     = params["id"]
		)

		if collection, err = collectionDialer(taskServiceURI, dialInfo.Database, TaskCollectionName); err == nil {
			defer collection.Close()
			log.Println("collection dialed successfully")
			err = collection.FindOne(taskID, task)
			log.Println("task search complete")
			statusCode = SuccessStatusResponseTaskByID
			response = task
		}
		r.JSON(statusCode, response)
	}
}
