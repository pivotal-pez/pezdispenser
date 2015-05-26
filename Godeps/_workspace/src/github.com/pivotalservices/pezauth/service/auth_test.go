package pezauth_test

import (
	"os"

	"github.com/go-martini/martini"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/pivotalservices/pezauth/service"
)

var _ = Describe("Authentication", func() {

	Describe("InitSession", func() {
		var (
			m *martini.ClassicMartini
		)

		BeforeEach(func() {
			setVcapApp("http://localhost:3000")
			setVcapServ()
			os.Setenv("PORT", "3000")
			m = martini.Classic()
		})

		Context("when InitSession is passed a classic martini", func() {
			It("Should setup the authentication middleware without panic", func() {
				Ω(func() {
					InitSession(m, &mockRedisCreds{})
				}).ShouldNot(Panic())
			})
		})

		Context("calling DomainCheck with a valid domain", func() {
			var (
				validDomain = "pivotal.io"
				validUser   = "testuser@pivotal.io"
			)
			setGetUserInfo(validDomain, validUser)

			It("Should have a valid statuscode and body", func() {
				mock := new(mockResponseWriter)
				DomainChecker(mock, new(mockTokens))
				Ω(mock.StatusCode).ShouldNot(Equal(FailureStatus))
				Ω(mock.Body).ShouldNot(Equal(AuthFailureResponse))
			})
		})

		Context("calling DomainCheck with a in-valid domain", func() {
			var (
				inValidDomain = "google.com"
				validUser     = "testuser@pivotal.io"
			)
			setGetUserInfo(inValidDomain, validUser)

			It("Should return true", func() {
				mock := new(mockResponseWriter)
				DomainChecker(mock, new(mockTokens))
				Ω(mock.StatusCode).Should(Equal(FailureStatus))
			})
		})
	})
})
