package pezauth_test

import (
	"fmt"
	"os"

	"github.com/go-martini/martini"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/pivotalservices/pezauth/service"
)

var _ = Describe("Routes", func() {
	Describe("InitRoutes", func() {
		var (
			m *martini.ClassicMartini
		)

		BeforeEach(func() {
			setVcapApp("http://localhost:3000")
			setVcapServ()
			os.Setenv("PORT", "3000")
			m = martini.Classic()
		})

		Context("calling InitSession with no enviornment variables set", func() {
			var (
				validDomain = "pivotal.io"
				validUser   = "testuser@pivotal.io"
			)

			setGetUserInfo(validDomain, validUser)

			BeforeEach(func() {
				os.Unsetenv("VCAP_APPLICATION")
				os.Unsetenv("VCAP_SERVICES")
			})

			It("Should panic", func() {
				Ω(func() {
					InitRoutes(m, new(mockDoer), new(mockMongo), new(mockHeritageClient))
				}).Should(Panic())
			})
		})

		Context("calling DomainCheck with a valid domain", func() {
			var (
				validDomain = "pivotal.io"
				validUser   = "testuser@pivotal.io"
			)
			setGetUserInfo(validDomain, validUser)

			Context("un-formatted domain", func() {
				BeforeEach(func() {
					setVcapApp("pivotal.io")
				})

				It("should format the domain in the config object", func() {
					control := fmt.Sprintf("https://%s/oauth2callback", validDomain)
					InitRoutes(m, new(mockDoer), new(mockMongo), new(mockHeritageClient))
					Ω(OauthConfig.RedirectURL).Should(Equal(control))
				})
			})

			Context("version formatted domain", func() {
				BeforeEach(func() {
					setVcapApp("pivotal-1919241972nwdighsd921h192t23t.io")
				})

				It("should format the domain in the config object", func() {
					control := fmt.Sprintf("https://%s/oauth2callback", validDomain)
					InitRoutes(m, new(mockDoer), new(mockMongo), new(mockHeritageClient))
					Ω(OauthConfig.RedirectURL).Should(Equal(control))
				})
			})
		})
	})
})
