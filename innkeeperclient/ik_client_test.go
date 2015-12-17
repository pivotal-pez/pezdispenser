package innkeeperclient_test

import (
	"net/http"

	. "github.com/pivotal-pez/pezdispenser/innkeeperclient"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
)

var _ = Describe("Given IkClient", func() {
	Describe("Given .GetStatus() method", func() {

		Context("When called with a valid requestid and that requestid has a status 'success'", func() {
			var (
				err               error
				res               *GetStatusResponse
				server            *ghttp.Server
				innkeeperUser     = "admin"
				innkeeperPassword = "pass"
			)
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyBasicAuth(innkeeperUser, innkeeperPassword),
						ghttp.RespondWith(http.StatusOK, `{
							"status": "success",
							"data": {
								"status": "ready",
								"credentials": {"name": "host-07-25", "oob_ip": "10.65.70.125", "oob_user": "user", "oob_pw": "xxxx"}
							},
							"message": "ok" 
						}`),
					),
				)
				ikClient := New(server.URL(), innkeeperUser, innkeeperPassword)
				res, err = ikClient.GetStatus("requestid")
			})
			AfterEach(func() {
				server.Close()
			})
			It("Then it should return an object containing current state from innkeeper", func() {
				Ω(err).ShouldNot(HaveOccurred())
				Ω(res).ShouldNot(BeNil())
				Ω(res.Data.Status).Should(Equal(StatusReady))
				Ω(res.Data.Credentials).ShouldNot(BeNil())
			})
		})
		Context("When called with a valid requestid that is 'running'", func() {
			var (
				err               error
				res               *GetStatusResponse
				server            *ghttp.Server
				innkeeperUser     = "admin"
				innkeeperPassword = "pass"
			)
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyBasicAuth(innkeeperUser, innkeeperPassword),
						ghttp.RespondWith(http.StatusOK, `{
							"status": "success",
							"data": {
								"status": "running",
								"credentials": {"name": "host-07-25", "oob_ip": "10.65.70.125", "oob_user": "user", "oob_pw": "xxxx"}
							},
							"message": "ok" 
						}`),
					),
				)
				ikClient := New(server.URL(), innkeeperUser, innkeeperPassword)
				res, err = ikClient.GetStatus("requestid")
			})
			AfterEach(func() {
				server.Close()
			})
			It("Then it should return an object containing current state from innkeeper", func() {
				Ω(err).ShouldNot(HaveOccurred())
				Ω(res).ShouldNot(BeNil())
				Ω(res.Data.Status).Should(Equal(StatusRunning))
			})
		})
	})
})
