package pezdispenser

import (
	"fmt"
	"os"

	"github.com/cloudfoundry-community/go-cfenv"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/pivotal-pez/pezdispenser/service/integrations"
)

const (
	//APIVersion1 - version 1 const
	APIVersion1 = "v1"
)

var (
	//URLBaseV1 - v1 url path base
	URLBaseV1 = fmt.Sprintf("/%s", APIVersion1)
)

//InitRoutes - initialize the mappings for controllers against valid routes
func InitRoutes(m *martini.ClassicMartini, keyCheckHandler martini.Handler, appEnv *cfenv.App) {
	taskServiceURI, taskServiceDatabase := getTaskBinding(appEnv)
	taskCollection := SetupDB(integrations.NewCollectionDialer, taskServiceURI, taskServiceDatabase, TaskCollectionName)
	m.Map(taskCollection)
	m.Use(render.Renderer())

	m.Group("/", func(r martini.Router) {
		r.Get("info", GetInfoController())
	})

	m.Group(URLBaseV1, func(r martini.Router) {
		r.Get("/task/:id", GetTaskByIDController())
		r.Post("/lease", PostLeaseController())
		r.Delete("/lease", DeleteLeaseController())
	}, keyCheckHandler)
}

func getTaskBinding(appEnv *cfenv.App) (taskServiceURI string, taskServiceDatabase string) {
	taskServiceName := os.Getenv("TASK_SERVICE_NAME")
	taskCredsURIName := os.Getenv("TASK_SERVICE_URI_NAME")
	taskCredsDBName := os.Getenv("TASK_SERVICE_DATABASE_NAME")

	if taskService, err := appEnv.Services.WithName(taskServiceName); err == nil {

		if taskServiceURI = taskService.Credentials[taskCredsURIName].(string); taskServiceURI == "" {
			panic(fmt.Sprint("we pulled an empty connection string %s from %v - %v", taskServiceURI, taskService, taskService.Credentials))
		}

		if taskServiceDatabase = taskService.Credentials[taskCredsDBName].(string); taskServiceDatabase == "" {
			panic(fmt.Sprint("we pulled an empty connection string %s from %v - %v", taskServiceDatabase, taskService, taskService.Credentials))
		}

	} else {
		panic(fmt.Sprint("Experienced an error trying to grab task service binding information:", err.Error()))
	}
	return
}
