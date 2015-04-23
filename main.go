package main

import (
	"os"

	cfenv "github.com/cloudfoundry-community/go-cfenv"
	"github.com/go-martini/martini"
	pez "github.com/pivotalservices/pezdispenser/service"
)

func main() {
	appEnv, _ := cfenv.Current()
	validatorServiceName := os.Getenv("UPS_PEZVALIDATOR_NAME")
	targetKeyName := os.Getenv("UPS_PEZVALIDATOR_TARGET")
	service, _ := appEnv.Services.WithName(validatorServiceName)
	validationTargetUrl := service.Credentials[targetKeyName]
	m := martini.Classic()
	pez.InitRoutes(m, validationTargetUrl)
	m.Run()
}
