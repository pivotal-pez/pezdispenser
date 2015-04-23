package pezauth_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("KeyGen", func() {
	var (
		username = "testuser@pivotal.io"
		details  = "{random:stuff}"
		guid     = "myfakekeyhash"
		keyhash  = fmt.Sprintf("%s:%s", username, guid)
		err      error
		response string
	)
	Context("Get function", func() {
		Context("calling get with a valid user arg", func() {
			BeforeEach(func() {
				k := getKeygen(false, keyhash, false)
				response, err = k.Get(username)
			})

			It("Should return an api key for that user", func() {
				Ω(response).Should(Equal(guid))
			})

			It("Should return a nil error", func() {
				Ω(err).Should(BeNil())
			})
		})

		Context("Get returns nil string", func() {
			It("Should not panic", func() {
				k := getKeygen(true, keyhash, true)
				Ω(func() {
					k.Get(username)
				}).ShouldNot(Panic())
			})
		})

		Context("calling get with a In-valid user arg", func() {
			BeforeEach(func() {
				k := getKeygen(true, keyhash, false)
				response, err = k.Get(username)
			})

			It("Should return an api key for that user", func() {
				Ω(response).ShouldNot(Equal(guid))
			})

			It("Should return a nil error", func() {
				Ω(err).ShouldNot(BeNil())
				Ω(err).Should(Equal(errDoerCallFailure))
			})
		})
	})

	Context("Create function", func() {
		Context("calling Create with a valid user arg", func() {
			BeforeEach(func() {
				k := getKeygen(false, keyhash, false)
				err = k.Create(username, details)
			})

			It("Should return a nil error", func() {
				Ω(err).Should(BeNil())
			})
		})

		Context("calling Create with a In-valid user arg", func() {
			BeforeEach(func() {
				k := getKeygen(true, keyhash, false)
				err = k.Create(username, details)
			})

			It("Should return a nil error", func() {
				Ω(err).ShouldNot(BeNil())
				Ω(err).Should(Equal(errDoerCallFailure))
			})
		})
	})

	Context("Delete function", func() {
		Context("calling Delete with a valid user arg", func() {
			BeforeEach(func() {
				k := getKeygen(false, keyhash, false)
				err = k.Delete(username)
			})

			It("Should return a nil error", func() {
				Ω(err).Should(BeNil())
			})
		})

		Context("calling Delete with a In-valid user arg", func() {
			BeforeEach(func() {
				k := getKeygen(true, keyhash, false)
				err = k.Create(username, details)
			})

			It("Should return a nil error", func() {
				Ω(err).ShouldNot(BeNil())
				Ω(err).Should(Equal(errDoerCallFailure))
			})
		})
	})
})
