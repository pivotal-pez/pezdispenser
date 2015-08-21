package pezdispenser

import (
	"fmt"
	"log"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/pivotal-pez/pezauth/keycheck"
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
func InitRoutes(m *martini.ClassicMartini, validationTargetUrl string) {
	keyCheckHandler := keycheck.NewAPIKeyCheckMiddleware(validationTargetUrl).Handler()
	m.Use(render.Renderer())

	m.Group("/", func(r martini.Router) {
		r.Get("info", func() string {
			return "the dispenser service will give you candy"
		})
	})

	m.Group(URLLeaseBaseV1, func(r martini.Router) {
		r.Get("/task/:id", func(params martini.Params, log *log.Logger, r render.Render) {
			taskID := params["id"]
			r.JSON(200, map[string]string{"taskID": taskID})
		})
	}, keyCheckHandler)
}
