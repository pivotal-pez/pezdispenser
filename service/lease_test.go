package pezdispenser_test

import (
	"bytes"
	"net/http"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pivotal-pez/pezdispenser/fakes"
	. "github.com/pivotal-pez/pezdispenser/service"
)

var _ = Describe("Lease", func() {
	Describe("NewLease", func() {
		Context("calling new lease", func() {
			It("should yield a lease valid lease", func() {
				lease := NewLease(new(fakes.FakeCollection))
				Ω(lease).ShouldNot(BeNil())
			})
		})
	})

	Describe(".Post()", func() {
		Context("called with an invalid request body", func() {
			var (
				lease   *Lease
				request *http.Request
			)
			BeforeEach(func() {
				request = new(http.Request)
				lease = NewLease(new(fakes.FakeCollection))
			})
			It("should return statuscode not found", func() {
				statusCode, _ := lease.Post(fakes.MockLogger, request)
				Ω(statusCode).Should(Equal(http.StatusNotFound))
			})
		})

		Context("called with a valid request body containing a lease", func() {
			var (
				lease   *Lease
				request *http.Request
			)
			BeforeEach(func() {
				request = new(http.Request)
				request.Body = fakes.FakeResponseBody{bytes.NewBufferString(`{"_id": "917397-292735-98293752935","inventory_id": "kaasd9sd9-98239h23h9-99h3ba993ba9h3ab","username": "someone","lease_duration": 14}`)}
				lease = NewLease(new(fakes.FakeCollection))
			})

			It("should return the lease object as the response", func() {
				statusCode, response := lease.Post(fakes.MockLogger, request)
				Ω(statusCode).Should(Equal(http.StatusCreated))
				Ω(response.(*Lease).Task.Status).Should(Equal(TaskStatusUnavailable))
				Ω(response.(*Lease).Task.Timestamp).ShouldNot(Equal(time.Time{}))
			})
		})
	})

	Describe(".InitFromHTTPRequest()", func() {
		Context("called with a valid request body containing a lease", func() {
			var (
				lease   *Lease
				request *http.Request
			)
			BeforeEach(func() {
				request = new(http.Request)
				request.Body = fakes.FakeResponseBody{bytes.NewBufferString(`{"_id": "917397-292735-98293752935","inventory_id": "kaasd9sd9-98239h23h9-99h3ba993ba9h3ab","username": "someone","lease_duration": 14}`)}
				lease = NewLease(new(fakes.FakeCollection))
			})

			It("should populate the lease's fields with the given lease", func() {
				err := lease.InitFromHTTPRequest(request)
				Ω(err).ShouldNot(HaveOccurred())
				Ω(lease.ID).Should(Equal("917397-292735-98293752935"))
				Ω(lease.InventoryID).Should(Equal("kaasd9sd9-98239h23h9-99h3ba993ba9h3ab"))
				Ω(lease.UserName).Should(Equal("someone"))
				Ω(lease.LeaseDuration).Should(Equal(float64(14)))
				Ω(lease.Task).Should(BeNil())
			})
		})
		Context("called with an invalid request body", func() {
			var (
				lease   *Lease
				request *http.Request
			)

			BeforeEach(func() {
				lease = NewLease(new(fakes.FakeCollection))
				request = new(http.Request)
			})

			It("should return a error", func() {
				err := lease.InitFromHTTPRequest(request)
				Ω(err).Should(HaveOccurred())
			})
		})
	})

	XDescribe(".Procurement()", func() {
		Context("when calling with a valid lease sku", func() {
			Context("of 2c.small", func() {
				It("should check the availablility, and allocate the resource", func() {

				})
			})
		})
	})

	XDescribe(".ReStock()", func() {
		Context("when calling with a valid lease sku", func() {
			Context("of 2c.small", func() {
				It("should check the state of the resource, begin restock process and update task info", func() {

				})
			})
		})
	})

	Describe(".SetTask()", func() {
		Context("calling with a valid task on an initialized lease", func() {
			var (
				lease       = NewLease(new(fakes.FakeCollection))
				controlTask = new(Task)
			)
			BeforeEach(func() {
				lease.SetTask(controlTask)
			})
			It("should set the task value", func() {
				Ω(lease.Task).Should(Equal(controlTask))
			})
		})
	})
})
