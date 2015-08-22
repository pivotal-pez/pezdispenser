package pezdispenser

import (
	"fmt"
	"os"

	"github.com/cloudfoundry-community/go-cfenv"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

//Constants to construct routes with
const (
	APIVersion1 = "v1"
)

//formatted strings based on constants, to be used in URLs
var (
	URLLeaseBaseV1 = fmt.Sprintf("/%s", APIVersion1)
)

//InitRoutes - initialize the mappings for controllers against valid routes
func InitRoutes(m *martini.ClassicMartini, keyCheckHandler martini.Handler, appEnv *cfenv.App) {
	taskServiceURI := getTaskBinding(appEnv)
	m.Use(render.Renderer())

	m.Group("/", func(r martini.Router) {
		r.Get("info", GetInfoController())
	})

	m.Group(URLLeaseBaseV1, func(r martini.Router) {
		r.Get("/task/:id", GetTaskByIdController(taskServiceURI))
	}, keyCheckHandler)
}

func getTaskBinding(appEnv *cfenv.App) (taskServiceURI string) {
	taskServiceName := os.Getenv("TASK_SERVICE_NAME")

	if taskService, err := appEnv.Services.WithName(taskServiceName); err == nil {
		taskServiceURI = fmt.Sprintf("%s", taskService.Credentials["TASK_SERVICE_URI_NAME"])

	} else {
		panic(fmt.Sprint("Experienced an error trying to grab task service binding information:", err.Error()))
	}
	return
}
