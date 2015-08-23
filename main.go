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
	fmt.Println("here")
	appEnv, err := cfenv.Current()
	fmt.Println("appEnv raw:", appEnv, err, os.Environ())
	validatorServiceName := os.Getenv("UPS_PEZVALIDATOR_NAME")
	targetKeyName := os.Getenv("UPS_PEZVALIDATOR_TARGET")
	fmt.Println("osGetEnv", validatorServiceName, targetKeyName)
	service, _ := appEnv.Services.WithName(validatorServiceName)
	fmt.Println("myservice", service)
	validationTargetURL := fmt.Sprintf("%s", service.Credentials[targetKeyName])
	fmt.Println("validationUrl:", validationTargetURL)
	m := martini.Classic()

	if appEnv, err := cfenv.Current(); err == nil {
		fmt.Println("setting up middleware keycheck")
		keyCheckHandler := keycheck.NewAPIKeyCheckMiddleware(validationTargetURL).Handler()
		fmt.Println("initroutes")
		pez.InitRoutes(m, keyCheckHandler, appEnv)
		fmt.Println("routes initialized")

	} else {
		panic(fmt.Sprint("Experienced an error trying to grab current cf environment:", err.Error()))
	}
	m.Run()
}
