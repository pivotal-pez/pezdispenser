package pezdispenser_test

import (
	"fmt"
	"os"

	"github.com/cloudfoundry-community/go-cfenv"
	"github.com/go-martini/martini"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/pivotal-pez/pezdispenser/service"
)

var _ = Describe("Routes", func() {
	Describe("InitRoutes", func() {
		var (
			m      *martini.ClassicMartini
			appEnv *cfenv.App
		)

		Context("when InitRoutes is passed valid arguments", func() {
			controlURI := "my-control-uri"
			controlTaskServiceName := "dispenser-task-service"

			BeforeEach(func() {
				m = martini.Classic()
				os.Setenv("TASK_SERVICE_NAME", controlTaskServiceName)
				appEnv, _ = cfenv.New(map[string]string{
					"VCAP_SERVICES":    fmt.Sprintf(vcapServicesFormatter, controlTaskServiceName, controlURI),
					"VCAP_APPLICATION": vcapApplicationFormatter,
				})
			})

			It("Should not result in panic", func() {
				Ω(func() {
					InitRoutes(m, fakeKeyCheck, appEnv)
				}).ShouldNot(Panic())
			})
		})

		Context("when InitRoutes is not given the proper appEnv", func() {
			BeforeEach(func() {
				m = martini.Classic()
				appEnv, _ = cfenv.New(map[string]string{
					"VCAP_SERVICES":    fmt.Sprintf(vcapServicesFormatter, "", ""),
					"VCAP_APPLICATION": vcapApplicationFormatter,
				})
			})

			It("Should panic and tell us what we are missing", func() {
				Ω(func() {
					InitRoutes(m, fakeKeyCheck, appEnv)
				}).Should(Panic())
			})
		})
	})
})
