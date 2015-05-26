package cloudfoundryclient_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/pivotalservices/pezdispenser/cloudfoundryclient"
)

var _ = Describe("CFClient", func() {
	Describe("AddOrg", func() {
		var (
			cfclient       CloudFoundryClient
			controlOrgGUID = "1e2bae2c-459e-4ad8-b1cb-ffc09d209b32"
			controlOrgName = "my-org-name"
		)

		Context("AddOrg called successfully", func() {

			BeforeEach(func() {
				mockDoer := &mockClientDoer{
					res: mockHttpResponse(mockSuccessOrgResponseBody, mockSuccessOrgStatusCode),
					err: nil,
				}
				mockRequest := &mockRequestDecorator{
					doer: mockDoer,
				}
				cfclient = NewCloudFoundryClient(mockRequest, new(mockLog))
			})

			It("should parse the response object without error", func() {
				guid, err := cfclient.AddOrg(controlOrgName)
				Ω(err).Should(BeNil())
				Ω(guid).Should(Equal(controlOrgGUID))
			})
		})

		Context("AddOrg unsuccessful response", func() {

			BeforeEach(func() {
				mockDoer := &mockClientDoer{
					res: mockHttpResponse(mockSuccessOrgResponseBody, (mockSuccessOrgStatusCode + 1)),
					err: nil,
				}
				mockRequest := &mockRequestDecorator{
					doer: mockDoer,
				}
				cfclient = NewCloudFoundryClient(mockRequest, new(mockLog))
			})

			It("should return an error", func() {
				guid, err := cfclient.AddOrg(controlOrgName)
				Ω(err).Should(Equal(ErrOrgCreateAPICallFailure))
				Ω(guid).Should(BeEmpty())
			})
		})
	})

	Describe("AddRole", func() {
		var cfclient CloudFoundryClient

		Context("AddRole called successfully", func() {

			BeforeEach(func() {
				mockDoer := &mockClientDoer{
					res: mockHttpResponse(mockSuccessRoleResponseBody, mockSuccessRoleStatusCode),
					err: nil,
				}
				mockRequest := &mockRequestDecorator{
					doer: mockDoer,
				}
				cfclient = NewCloudFoundryClient(mockRequest, new(mockLog))
			})

			It("should parse the response object without error", func() {
				err := cfclient.AddRole(OrgEndpoint, "target-guid-12345", RoleTypeManager, "user-guid-12345")
				Ω(err).Should(BeNil())
			})
		})

		Context("AddRole unsuccessful response", func() {

			BeforeEach(func() {
				mockDoer := &mockClientDoer{
					res: mockHttpResponse(mockSuccessRoleResponseBody, (mockSuccessRoleStatusCode + 1)),
					err: nil,
				}
				mockRequest := &mockRequestDecorator{
					doer: mockDoer,
				}
				cfclient = NewCloudFoundryClient(mockRequest, new(mockLog))
			})

			It("should return an error", func() {
				err := cfclient.AddRole(OrgEndpoint, "target-guid-12345", RoleTypeManager, "user-guid-12345")
				Ω(err).Should(Equal(ErrFailedStatusCode))
			})
		})
	})

	Describe("QueryUserGUID", func() {
		var cfclient CloudFoundryClient

		Context("QueryUserGUID called successfully", func() {

			BeforeEach(func() {
				mockDoer := &mockClientDoer{
					res: mockHttpResponse(mockSuccessUserResponseBody, mockSuccessUserStatusCode),
					err: nil,
				}
				mockRequest := &mockRequestDecorator{
					doer: mockDoer,
				}
				cfclient = NewCloudFoundryClient(mockRequest, new(mockLog))
			})

			It("should parse the response object without error", func() {
				controlUser := "testuser"
				controlUID := "123456"
				guid, err := cfclient.QueryUserGUID(controlUser)
				Ω(guid).Should(Equal(controlUID))
				Ω(err).Should(BeNil())
			})
		})

		Context("QueryUserGUID called w/ invalid user", func() {

			BeforeEach(func() {
				mockDoer := &mockClientDoer{
					res: mockHttpResponse(mockSuccessUserResponseBody, mockSuccessUserStatusCode),
					err: nil,
				}
				mockRequest := &mockRequestDecorator{
					doer: mockDoer,
				}
				cfclient = NewCloudFoundryClient(mockRequest, new(mockLog))
			})

			It("should parse the return a user not found error", func() {
				controlUser := "invalid-user"
				guid, err := cfclient.QueryUserGUID(controlUser)
				Ω(guid).Should(BeEmpty())
				Ω(err).Should(Equal(ErrNoUserFound))
			})
		})

		Context("QueryUserGUID call failed", func() {

			BeforeEach(func() {
				mockDoer := &mockClientDoer{
					res: mockHttpResponse(mockSuccessUserResponseBody, (mockSuccessUserStatusCode + 1)),
					err: nil,
				}
				mockRequest := &mockRequestDecorator{
					doer: mockDoer,
				}
				cfclient = NewCloudFoundryClient(mockRequest, new(mockLog))
			})

			It("should parse the return a user not found error", func() {
				controlUser := "testuser"
				guid, err := cfclient.QueryUserGUID(controlUser)
				Ω(guid).Should(BeEmpty())
				Ω(err).Should(Equal(ErrFailedStatusCode))
			})
		})
	})

	Describe("QueryAPIInfo", func() {
		var cfclient CloudFoundryClient

		Context("QueryAPIInfo called successfully", func() {

			BeforeEach(func() {
				mockDoer := &mockClientDoer{
					res: mockHttpResponse(mockSuccessInfoResponseBody, mockSuccessInfoStatusCode),
					err: nil,
				}
				mockRequest := &mockRequestDecorator{
					doer: mockDoer,
				}
				cfclient = NewCloudFoundryClient(mockRequest, new(mockLog))
			})

			It("should parse the response object without error", func() {
				info, err := cfclient.QueryAPIInfo()
				Ω(info.LoggingEndpoint).ShouldNot(BeEmpty())
				Ω(info.AuthorizationEndpoint).ShouldNot(BeEmpty())
				Ω(info.TokenEndpoint).ShouldNot(BeEmpty())
				Ω(err).Should(BeNil())
			})
		})

		Context("QueryAPIInfo called with failure", func() {

			BeforeEach(func() {
				mockDoer := &mockClientDoer{
					res: mockHttpResponse(mockSuccessInfoResponseBody, (mockSuccessInfoStatusCode + 1)),
					err: nil,
				}
				mockRequest := &mockRequestDecorator{
					doer: mockDoer,
				}
				cfclient = NewCloudFoundryClient(mockRequest, new(mockLog))
			})

			It("should parse the response object without error", func() {
				info, err := cfclient.QueryAPIInfo()
				Ω(info.LoggingEndpoint).Should(BeEmpty())
				Ω(info.AuthorizationEndpoint).Should(BeEmpty())
				Ω(info.TokenEndpoint).Should(BeEmpty())
				Ω(err).ShouldNot(BeNil())
			})
		})
	})
})
