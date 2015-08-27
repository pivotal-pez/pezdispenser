package pezdispenser

import (
	"log"
	"net/http"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/pivotal-pez/pezdispenser/service/integrations"
	"github.com/pivotal-pez/pezdispenser/taskmanager"
)

//GetTaskByIDController - this is the controller to handle a get task call
func GetTaskByIDController(taskServiceURI string, collectionDialer integrations.CollectionDialer) martini.Handler {
	taskCollection := setupDB(collectionDialer, taskServiceURI, TaskCollectionName)

	return func(params martini.Params, logger *log.Logger, r render.Render) {
		var (
			err        error
			response   interface{}
			statusCode = http.StatusNotFound
			task       = new(taskmanager.Task)
			taskID     = params["id"]
		)
		taskCollection.Wake()
		logger.Println("collection dialed successfully")

		if err = taskCollection.FindOne(taskID, task); err == nil {
			logger.Println("task search complete")
			statusCode = http.StatusOK
			response = task

		} else {
			response = map[string]string{"error": err.Error()}
		}
		r.JSON(statusCode, response)
	}
}
