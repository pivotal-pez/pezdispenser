package main

import (
	"fmt"
	"os"

	cfenv "github.com/cloudfoundry-community/go-cfenv"
	"github.com/go-martini/martini"
	"github.com/pivotal-pez/pezauth/keycheck"
	pez "github.com/pivotal-pez/pezdispenser/service"
)

func main() {
	appEnv, _ := cfenv.Current()
	validatorServiceName := os.Getenv("UPS_PEZVALIDATOR_NAME")
	targetKeyName := os.Getenv("UPS_PEZVALIDATOR_TARGET")
	service, _ := appEnv.Services.WithName(validatorServiceName)
	validationTargetURL := fmt.Sprintf("%s", service.Credentials[targetKeyName])
	m := martini.Classic()

	if appEnv, err := cfenv.Current(); err == nil {
		keyCheckHandler := keycheck.NewAPIKeyCheckMiddleware(validationTargetURL).Handler()
		pez.InitRoutes(m, keyCheckHandler, appEnv)

	} else {
		panic(fmt.Sprint("Experienced an error trying to grab current cf environment:", err.Error()))
	}
	m.Run()
}
