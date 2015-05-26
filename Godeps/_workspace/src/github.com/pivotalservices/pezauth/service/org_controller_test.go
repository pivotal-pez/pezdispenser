package pezauth_test

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/fatih/structs"
	"github.com/go-martini/martini"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/pivotalservices/pezauth/service"
)

var _ = Describe("NewOrgController", func() {

	Describe("Put Control handler", func() {
		var (
			fakeName             = "testuser"
			fakeUser             = fmt.Sprintf("%s@pivotal.io", fakeName)
			fakeOrg              = fmt.Sprintf("pivot-%s", fakeName)
			render               *mockRenderer
			testLogger           = log.New(os.Stdout, "testLogger", 0)
			controlGUID          = "7cc7b460-a626-4189-8c93-726373e7a492"
			successOrgCreateBody = `{
  "metadata": {
    "guid": "7cc7b460-a626-4189-8c93-726373e7a492",
    "url": "/v2/organizations/7cc7b460-a626-4189-8c93-726373e7a492",
    "created_at": "2015-04-15T00:55:51Z",
    "updated_at": null
  },
  "entity": {
    "name": "my-org-name",
    "billing_enabled": false,
    "quota_definition_guid": "a7196e89-e2a9-472e-8930-56bdedbf0308",
    "status": "active",
    "quota_definition_url": "/v2/quota_definitions/a7196e89-e2a9-472e-8930-56bdedbf0308",
    "spaces_url": "/v2/organizations/7cc7b460-a626-4189-8c93-726373e7a492/spaces",
    "domains_url": "/v2/organizations/7cc7b460-a626-4189-8c93-726373e7a492/domains",
    "private_domains_url": "/v2/organizations/7cc7b460-a626-4189-8c93-726373e7a492/private_domains",
    "users_url": "/v2/organizations/7cc7b460-a626-4189-8c93-726373e7a492/users",
    "managers_url": "/v2/organizations/7cc7b460-a626-4189-8c93-726373e7a492/managers",
    "billing_managers_url": "/v2/organizations/7cc7b460-a626-4189-8c93-726373e7a492/billing_managers",
    "auditors_url": "/v2/organizations/7cc7b460-a626-4189-8c93-726373e7a492/auditors",
    "app_events_url": "/v2/organizations/7cc7b460-a626-4189-8c93-726373e7a492/app_events",
    "space_quota_definitions_url": "/v2/organizations/7cc7b460-a626-4189-8c93-726373e7a492/space_quota_definitions"
  }
}`
		)
		setGetUserInfo("pivotal.io", fakeUser)

		BeforeEach(func() {
			render = new(mockRenderer)
		})

		Context("when email is not in the system", func() {
			var orgPut OrgPutHandler
			var oldNewOrg = NewOrg
			tokens := &mockTokens{}
			result := PivotOrg{
				Email:   fakeUser,
				OrgName: fakeOrg,
				OrgGUID: controlGUID,
			}
			controlResponse := Response{Payload: structs.Map(result)}

			BeforeEach(func() {
				NewOrg = getMockNewOrg(nil, &result, &result, ErrNoMatchInStore, nil, nil)
				orgPut = NewOrgController(&mockPersistence{
					err:    ErrNoMatchInStore,
					result: "",
				}, &mockHeritageClient{
					res: &http.Response{
						StatusCode: 201,
						Body:       nopCloser{bytes.NewBufferString(successOrgCreateBody)},
					},
				}).Put().(OrgPutHandler)
			})

			AfterEach(func() {
				NewOrg = oldNewOrg
			})

			It("should create a new org record", func() {
				orgPut(martini.Params{UserParam: fakeUser}, testLogger, render, tokens)
				Ω(render.StatusCode).Should(Equal(SuccessStatus))
			})

			It("should create a new org record", func() {
				orgPut(martini.Params{UserParam: fakeUser}, testLogger, render, tokens)
				Ω(render.ResponseObject).Should(Equal(controlResponse))
			})
		})

		Context("when org create fails", func() {
			tokens := &mockTokens{}
			result := PivotOrg{
				Email:   fakeUser,
				OrgName: fakeOrg,
			}
			controlResponse := Response{Payload: structs.Map(result)}
			var orgPut OrgPutHandler = NewOrgController(&mockPersistence{
				err:    ErrNoMatchInStore,
				result: nil,
			}, &mockHeritageClient{
				res: &http.Response{
					StatusCode: 403,
					Body:       nopCloser{bytes.NewBufferString(`random test response`)},
				},
			}).Put().(OrgPutHandler)

			It("should return an error response", func() {
				orgPut(martini.Params{UserParam: fakeUser}, testLogger, render, tokens)
				Ω(render.StatusCode).Should(Equal(FailureStatus))
				Ω(render.ResponseObject).ShouldNot(Equal(controlResponse))
			})
		})
	})

	Describe("Get Control handler", func() {
		Context("calling controller with a bad user token combo", func() {
			var (
				fakeName   = "testuser"
				fakeUser   = fmt.Sprintf("%s@pivotal.io", fakeName)
				badUser    = fmt.Sprintf("%s@pivotal.io", "baduser")
				render     *mockRenderer
				testLogger = log.New(os.Stdout, "testLogger", 0)
			)
			setGetUserInfo("pivotal.io", badUser)

			BeforeEach(func() {
				render = new(mockRenderer)
			})

			Context("with a user that has no match in the system", func() {
				tokens := &mockTokens{}
				result := new(PivotOrg)
				controlResponse := Response{ErrorMsg: ErrCantCallAcrossUsers.Error()}
				var orgGet OrgGetHandler = NewOrgController(&mockPersistence{
					err:    ErrCantCallAcrossUsers,
					result: result,
				}, new(mockHeritageClient)).Get().(OrgGetHandler)

				It("should return an error and a fail status", func() {
					orgGet(martini.Params{UserParam: fakeUser}, testLogger, render, tokens)
					Ω(render.StatusCode).Should(Equal(FailureStatus))
					Ω(render.ResponseObject).Should(Equal(controlResponse))
				})
			})

			Context("with a user that has an org", func() {
				tokens := &mockTokens{}
				result := new(PivotOrg)
				controlResponse := Response{ErrorMsg: ErrCantCallAcrossUsers.Error()}
				var orgGet OrgGetHandler = NewOrgController(&mockPersistence{
					err:    ErrCantCallAcrossUsers,
					result: result,
				}, new(mockHeritageClient)).Get().(OrgGetHandler)

				It("should return a error object to the renderer", func() {
					orgGet(martini.Params{UserParam: fakeUser}, testLogger, render, tokens)
					Ω(render.StatusCode).Should(Equal(FailureStatus))
					Ω(render.ResponseObject).Should(Equal(controlResponse))
				})
			})
		})

		Context("calling controller", func() {
			var (
				fakeName   = "testuser"
				fakeUser   = fmt.Sprintf("%s@pivotal.io", fakeName)
				fakeOrg    = fmt.Sprintf("pivot-%s", fakeName)
				render     *mockRenderer
				testLogger = log.New(os.Stdout, "testLogger", 0)
			)
			setGetUserInfo("pivotal.io", fakeUser)

			BeforeEach(func() {
				render = new(mockRenderer)
			})

			Context("with a user that has no match in the system", func() {
				tokens := &mockTokens{}
				result := PivotOrg{
					Email:   fakeUser,
					OrgName: fakeOrg,
				}
				controlResponse := Response{ErrorMsg: ErrNoMatchInStore.Error()}
				var orgGet OrgGetHandler = NewOrgController(&mockPersistence{
					err:    ErrNoMatchInStore,
					result: result,
				}, new(mockHeritageClient)).Get().(OrgGetHandler)

				It("should return an error and a fail status", func() {
					orgGet(martini.Params{UserParam: fakeUser}, testLogger, render, tokens)
					Ω(render.StatusCode).Should(Equal(FailureStatus))
					Ω(render.ResponseObject).Should(Equal(controlResponse))
				})
			})

			Context("with a user that has an org", func() {
				tokens := &mockTokens{}
				result := PivotOrg{
					Email:   fakeUser,
					OrgName: fakeOrg,
				}
				controlResponse := Response{Payload: structs.Map(result)}
				var orgGet OrgGetHandler = NewOrgController(&mockPersistence{
					err:    nil,
					result: result,
				}, new(mockHeritageClient)).Get().(OrgGetHandler)

				It("should return a user object to the renderer", func() {
					orgGet(martini.Params{UserParam: fakeUser}, testLogger, render, tokens)
					Ω(render.StatusCode).Should(Equal(SuccessStatus))
					Ω(render.ResponseObject).Should(Equal(controlResponse))
				})
			})
		})
	})
})
