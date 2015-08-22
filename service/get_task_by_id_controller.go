package pezdispenser

import (
	"log"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

//GetTaskByIdController - this is the controller to handle a get task call
func GetTaskByIdController(taskServiceURI string) martini.Handler {
	return func(params martini.Params, log *log.Logger, r render.Render) {
		taskID := params["id"]
		log.Println(taskServiceURI)
		r.JSON(200, map[string]string{"taskID": taskID})
	}
}
