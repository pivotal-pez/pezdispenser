package vcloudclient_test

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/pivotal-pez/pezdispenser/vcloudclient"
)

var _ = Describe("VCloud Client", func() {
	Describe("VCDClient", func() {
		Describe(".UnDeployVApp()", func() {
			var (
				vcdClient    *VCDClient
				controlToken = "xxxxxxxxxxxxxxxxxedw8d8sdb9sdb9sdbsd9sdbsdb"
			)
			Context("when the client call fails", func() {
				var randomClientError = errors.New("random client error")
				BeforeEach(func() {
					client := new(fakeHttpClient)
					client.Err = randomClientError
					vcdClient = NewVCDClient(client, "")
					vcdClient.Token = controlToken
				})
				It("should return the client error", func() {
					_, err := vcdClient.UnDeployVApp("vappid")
					Ω(err).Should(HaveOccurred())
					Ω(err).Should(Equal(randomClientError))
				})
			})

			Context("when the REST call returns a non successful statuscode", func() {
				BeforeEach(func() {
					xmlResponse := fmt.Sprintf(TaskResponseFormatter, "queued")
					client := new(fakeHttpClient)
					client.Response = new(http.Response)
					client.Response.StatusCode = (DeleteVappSuccessStatusCode + 201)
					client.Response.Body = nopCloser{bytes.NewBufferString(xmlResponse)}
					vcdClient = NewVCDClient(client, "")
					vcdClient.Token = controlToken
				})

				It("should return a failure status error", func() {
					_, err := vcdClient.UnDeployVApp("vappid")
					Ω(err).Should(HaveOccurred())
					Ω(err).Should(Equal(ErrTaskResponseParseFailed))
				})
			})

			Context("when called with a valid vapp id", func() {

				BeforeEach(func() {
					xmlResponse := fmt.Sprintf(TaskResponseFormatter, "queued")
					client := new(fakeHttpClient)
					client.Response = new(http.Response)
					client.Response.StatusCode = DeleteVappSuccessStatusCode
					client.Response.Body = nopCloser{bytes.NewBufferString(xmlResponse)}
					vcdClient = NewVCDClient(client, "")
					vcdClient.Token = controlToken
				})

				It("should make the call to the delete vapp api endpoint and return a task to monitor the deletion of the vapp", func() {
					task, err := vcdClient.UnDeployVApp("vappid")
					Ω(err).ShouldNot(HaveOccurred())
					Ω(task.Status).ShouldNot(BeEmpty())
				})
			})
		})

		Describe(".DeleteVApp()", func() {
			var (
				vcdClient    *VCDClient
				controlToken = "xxxxxxxxxxxxxxxxxedw8d8sdb9sdb9sdbsd9sdbsdb"
			)
			Context("when the client call fails", func() {
				var randomClientError = errors.New("random client error")
				BeforeEach(func() {
					client := new(fakeHttpClient)
					client.Err = randomClientError
					vcdClient = NewVCDClient(client, "")
					vcdClient.Token = controlToken
				})
				It("should return the client error", func() {
					_, err := vcdClient.DeleteVApp("vappid")
					Ω(err).Should(HaveOccurred())
					Ω(err).Should(Equal(randomClientError))
				})
			})

			Context("when the REST call returns a non successful statuscode", func() {
				BeforeEach(func() {
					xmlResponse := fmt.Sprintf(TaskResponseFormatter, "queued")
					client := new(fakeHttpClient)
					client.Response = new(http.Response)
					client.Response.StatusCode = (DeleteVappSuccessStatusCode + 201)
					client.Response.Body = nopCloser{bytes.NewBufferString(xmlResponse)}
					vcdClient = NewVCDClient(client, "")
					vcdClient.Token = controlToken
				})

				It("should return a failure status error", func() {
					_, err := vcdClient.DeleteVApp("vappid")
					Ω(err).Should(HaveOccurred())
					Ω(err).Should(Equal(ErrTaskResponseParseFailed))
				})
			})

			Context("when called with a valid vapp id", func() {

				BeforeEach(func() {
					xmlResponse := fmt.Sprintf(TaskResponseFormatter, "queued")
					client := new(fakeHttpClient)
					client.Response = new(http.Response)
					client.Response.StatusCode = DeleteVappSuccessStatusCode
					client.Response.Body = nopCloser{bytes.NewBufferString(xmlResponse)}
					vcdClient = NewVCDClient(client, "")
					vcdClient.Token = controlToken
				})

				It("should make the call to the delete vapp api endpoint and return a task to monitor the deletion of the vapp", func() {
					task, err := vcdClient.DeleteVApp("vappid")
					Ω(err).ShouldNot(HaveOccurred())
					Ω(task.Status).ShouldNot(BeEmpty())
				})
			})
		})
		Describe(".PollTaskURL()", func() {
			var (
				vcdClient    *VCDClient
				controlToken = "xxxxxxxxxxxxxxxxxedw8d8sdb9sdb9sdbsd9sdbsdb"
			)
			Context("when the client call fails", func() {
				var randomClientError = errors.New("random client error")
				BeforeEach(func() {
					client := new(fakeHttpClient)
					client.Err = randomClientError
					vcdClient = NewVCDClient(client, "")
					vcdClient.Token = controlToken
				})
				It("should return the client error", func() {
					_, err := vcdClient.PollTaskURL("vappid")
					Ω(err).Should(HaveOccurred())
					Ω(err).Should(Equal(randomClientError))
				})
			})

			Context("when the REST call returns a non successful statuscode", func() {
				BeforeEach(func() {
					xmlResponse := fmt.Sprintf(TaskResponseFormatter, "queued")
					client := new(fakeHttpClient)
					client.Response = new(http.Response)
					client.Response.StatusCode = (TaskPollSuccessStatusCode + 201)
					client.Response.Body = nopCloser{bytes.NewBufferString(xmlResponse)}
					vcdClient = NewVCDClient(client, "")
					vcdClient.Token = controlToken
				})

				It("should return a failure status error", func() {
					_, err := vcdClient.PollTaskURL("vappid")
					Ω(err).Should(HaveOccurred())
					Ω(err).Should(Equal(ErrTaskResponseParseFailed))
				})
			})

			Context("when called with a valid vapp id", func() {

				BeforeEach(func() {
					xmlResponse := fmt.Sprintf(TaskResponseFormatter, "queued")
					client := new(fakeHttpClient)
					client.Response = new(http.Response)
					client.Response.StatusCode = TaskPollSuccessStatusCode
					client.Response.Body = nopCloser{bytes.NewBufferString(xmlResponse)}
					vcdClient = NewVCDClient(client, "")
					vcdClient.Token = controlToken
				})

				It("should make the call to the task api endpoint and return a task to monitor its status", func() {
					task, err := vcdClient.PollTaskURL("vappid")
					Ω(err).ShouldNot(HaveOccurred())
					Ω(task.Status).ShouldNot(BeEmpty())
				})
			})
		})

		Describe(".DeployVApp()", func() {
			var (
				vcdClient       *VCDClient
				controlToken    = "xxxxxxxxxxxxxxxxxedw8d8sdb9sdb9sdbsd9sdbsdb"
				controlSlotName = "PCFaaS-Slot-10"
				controlVcdHref  = "https://sandbox.pez.pivotal.io/api/vdc/59b61466-fad9-49b4-a355-2467d311da78"
				controlHref     = "https://sandbox.pez.pivotal.io/api/vAppTemplate/vappTemplate-8b761107-eddc-430c-8aba-3cdf900e9812"
			)
			Context("when a call to the rest api fails", func() {
				BeforeEach(func() {
					client := new(fakeHttpClient)
					client.Response = new(http.Response)
					client.Response.StatusCode = (DeployVappSuccessStatusCode + 201)
					vcdClient = NewVCDClient(client, "")
					vcdClient.Token = controlToken
				})

				It("should yield an error showing the failure", func() {
					_, err := vcdClient.DeployVApp(controlSlotName, controlHref, controlVcdHref)
					Ω(err).Should(HaveOccurred())
				})
			})

			Context("when called with valid templatename, templatehref & vcdhref", func() {
				controlTaskHref := "https://sampleurl.com"

				BeforeEach(func() {
					client := new(fakeHttpClient)
					client.Response = new(http.Response)
					client.Response.StatusCode = DeployVappSuccessStatusCode
					fixtureData, _ := ioutil.ReadFile("fixtures/deploy_vapp_response.xml")
					client.Response.Body = nopCloser{bytes.NewBuffer(fixtureData)}
					vcdClient = NewVCDClient(client, "")
					vcdClient.Token = controlToken
				})

				It("should not yield an error", func() {
					_, err := vcdClient.DeployVApp(controlSlotName, controlHref, controlVcdHref)
					Ω(err).ShouldNot(HaveOccurred())
				})

				It("should return a valid vapp object for the deployment call", func() {
					vapp, _ := vcdClient.DeployVApp(controlSlotName, controlHref, controlVcdHref)
					Ω(vapp.Tasks.Task.Href).Should(Equal(controlTaskHref))
				})
			})
		})

		Describe(".QueryTemplate()", func() {
			var (
				vcdClient       *VCDClient
				controlToken    = "xxxxxxxxxxxxxxxxxedw8d8sdb9sdb9sdbsd9sdbsdb"
				controlSlotName = "PCFaaS-Slot-10"
			)

			Context("when query call response has status other than 200", func() {
				BeforeEach(func() {
					client := new(fakeHttpClient)
					client.Response = new(http.Response)
					client.Response.StatusCode = (QuerySuccessStatusCode + 201)
					vcdClient = NewVCDClient(client, "")
					vcdClient.Token = controlToken
				})

				It("Should return query failed error", func() {
					_, err := vcdClient.QueryTemplate(controlSlotName)
					Ω(err).Should(HaveOccurred())
					Ω(err).Should(Equal(ErrFailedQuery))
				})
			})

			Context("when given a valid template name", func() {
				BeforeEach(func() {
					client := new(fakeHttpClient)
					client.Response = new(http.Response)
					client.Response.StatusCode = QuerySuccessStatusCode
					client.Response.Header = http.Header{}
					client.Response.Header.Set(VCloudTokenHeaderName, controlToken)
					fixtureData, _ := ioutil.ReadFile("fixtures/template_query_response.xml")
					client.Response.Body = nopCloser{bytes.NewBuffer(fixtureData)}
					vcdClient = NewVCDClient(client, "")
					vcdClient.Token = controlToken
				})

				It("Should return a vapptemplate object for the matching template", func() {
					template, err := vcdClient.QueryTemplate(controlSlotName)
					Ω(err).ShouldNot(HaveOccurred())
					Ω(template.Name).Should(Equal(controlSlotName))
					Ω(template.Vdc).Should(Equal("https://sandbox.pez.pivotal.io/api/vdc/59b61466-fad9-49b4-a355-2467d311da78"))
					Ω(template.Href).Should(Equal("https://sandbox.pez.pivotal.io/api/vAppTemplate/vappTemplate-8b761107-eddc-430c-8aba-3cdf900e9812"))
				})
			})
		})

		Describe(".AuthDecorate()", func() {
			var (
				vcdClient    *VCDClient
				controlToken = "xxxxxxxxxxxxxxxxxedw8d8sdb9sdb9sdbsd9sdbsdb"
			)
			Context("given an *http.Request object", func() {
				BeforeEach(func() {
					client := new(fakeHttpClient)
					vcdClient = NewVCDClient(client, "")
				})

				It("should add the proper authentication token to the header of the given request", func() {
					vcdClient.Token = controlToken
					req := new(http.Request)
					vcdClient.AuthDecorate(req)
					token := req.Header.Get(VCloudTokenHeaderName)
					Ω(token).Should(Equal(controlToken))
				})

				It("should return error if there is no token available", func() {
					req := new(http.Request)
					err := vcdClient.AuthDecorate(req)
					Ω(err).Should(HaveOccurred())
					Ω(err).Should(Equal(ErrNoTokenToApply))
				})
			})
		})
		Describe(".Auth()", func() {
			var (
				vcdClient    *VCDClient
				controlToken = "xxxxxxxxxxxxxxxxxedw8d8sdb9sdb9sdbsd9sdbsdb"
			)

			Context("given valid user/pass/uri", func() {
				BeforeEach(func() {
					client := new(fakeHttpClient)
					client.Response = new(http.Response)
					client.Response.StatusCode = AuthSuccessStatusCode
					client.Response.Header = http.Header{}
					client.Response.Header.Set(VCloudTokenHeaderName, controlToken)
					vcdClient = NewVCDClient(client, "")
				})

				It("should set us a valid auth token", func() {
					err := vcdClient.Auth("", "")
					token := vcdClient.Token
					Ω(err).ShouldNot(HaveOccurred())
					Ω(token).ShouldNot(BeEmpty())
					Ω(token).Should(Equal(controlToken))
				})
			})

			Context("given the api does not authenticate against our credentials", func() {
				BeforeEach(func() {
					client := new(fakeHttpClient)
					client.Response = new(http.Response)
					client.Response.StatusCode = (AuthSuccessStatusCode + 201)
					client.Response.Header = http.Header{}
					client.Response.Header.Set(VCloudTokenHeaderName, controlToken)
					vcdClient = NewVCDClient(client, "")
				})

				It("should return the proper error", func() {
					err := vcdClient.Auth("", "")
					Ω(err).Should(HaveOccurred())
					Ω(err).Should(Equal(ErrAuthFailure))
				})

				It("should not set a token", func() {
					vcdClient.Auth("", "")
					token := vcdClient.Token
					Ω(token).Should(BeEmpty())
				})
			})

			Context("given an authentication call returns error", func() {
				BeforeEach(func() {
					client := new(fakeHttpClient)
					client.Err = errors.New("random connection error")
					client.Response = new(http.Response)
					client.Response.StatusCode = (AuthSuccessStatusCode + 300)
					vcdClient = NewVCDClient(client, "")
				})

				It("should pass through the error from the client connection", func() {
					err := vcdClient.Auth("", "")
					token := vcdClient.Token
					Ω(err).Should(HaveOccurred())
					Ω(err).ShouldNot(Equal(ErrAuthFailure))
					Ω(token).Should(BeEmpty())
				})
			})
		})
	})
})
