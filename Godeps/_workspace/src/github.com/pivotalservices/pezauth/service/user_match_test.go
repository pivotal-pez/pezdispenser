package pezauth_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/pivotalservices/pezauth/service"
)

var _ = Describe("NewUserMatch", func() {
	Context("when given a matching username + token combo", func() {
		var (
			userMatch *UserMatch
			username  = "testuser@pivotal.io"
			userInfo  = map[string]interface{}{
				"domain": "pivotal.io",
				"emails": []interface{}{
					map[string]interface{}{
						"value": "garbage",
					},
					map[string]interface{}{
						"value": username,
					},
				},
			}
		)

		BeforeEach(func() {
			userMatch = NewUserMatch().
				UserInfo(userInfo).
				UserName(username)
		})

		It("should return a nil error and run the on success method", func() {
			successCnt := 0
			failCnt := 0

			err := userMatch.OnSuccess(func() {
				successCnt++
			}).
				OnFailure(func() {
				failCnt++
			}).Run()
			Ω(err).Should(BeNil())
			Ω(failCnt).Should(Equal(0))
			Ω(successCnt).Should(Equal(1))
		})
	})

	Context("when given a NON-matching username + token combo", func() {
		var (
			userMatch *UserMatch
			username  = "testuser@pivotal.io"
			userInfo  = map[string]interface{}{
				"domain": "pivotal.io",
				"emails": []interface{}{
					map[string]interface{}{
						"value": "garbage",
					},
					map[string]interface{}{
						"value": "wrongUser@pivotal.io",
					},
				},
			}
		)

		BeforeEach(func() {
			userMatch = NewUserMatch().
				UserInfo(userInfo).
				UserName(username)
		})

		It("should return a nil error and run the on success method", func() {
			successCnt := 0
			failCnt := 0

			err := userMatch.OnSuccess(func() {
				successCnt++
			}).
				OnFailure(func() {
				failCnt++
			}).Run()
			Ω(err).Should(Equal(ErrNotValidActionForUser))
			Ω(failCnt).Should(Equal(1))
			Ω(successCnt).Should(Equal(0))
		})
	})
})
