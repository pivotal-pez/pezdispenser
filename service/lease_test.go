package pezdispenser_test

import (
	"bytes"
	"net/http"
	"time"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pivotal-pez/pezdispenser/fakes"
	. "github.com/pivotal-pez/pezdispenser/service"
	"github.com/pivotal-pez/pezdispenser/taskmanager"
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

	Describe(".InventoryAvailable()", func() {
		Context("when we can find the record but taskStatus is != available", func() {
			var (
				lease *Lease
			)
			BeforeEach(func() {
				lease = NewLease(new(fakes.FakeCollection))
			})
			It("should indicate a unavailable status of the inventory item", func() {
				yesno := lease.InventoryAvailable()
				Ω(yesno).Should(BeFalse())
			})
		})

		Context("when we can find the record and taskStatus is == available", func() {
			var (
				lease *Lease
			)
			BeforeEach(func() {
				col := new(fakes.FakeCollection)
				col.ControlTask = taskmanager.Task{
					Status: TaskStatusAvailable,
				}
				lease = NewLease(col)
			})
			It("should indicate a available status of the inventory item", func() {
				yesno := lease.InventoryAvailable()
				Ω(yesno).Should(BeTrue())
			})
		})

		Context("when we can not find the record", func() {
			var (
				lease *Lease
			)
			BeforeEach(func() {
				col := new(fakes.FakeCollection)
				col.ErrControl = mgo.ErrNotFound
				lease = NewLease(col)
				lease.InventoryID = bson.NewObjectId().Hex()
			})
			It("should indicate a available status of the inventory item", func() {
				yesno := lease.InventoryAvailable()
				Ω(yesno).Should(BeTrue())
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

	Describe(".Procurement()", func() {
		Context("when calling with a valid lease (default)", func() {
			var (
				lease   *Lease
				request *http.Request
			)
			BeforeEach(func() {
				request = new(http.Request)
				request.Body = fakes.FakeResponseBody{bytes.NewBufferString(`{"_id": "917397-292735-98293752935","inventory_id": "kaasd9sd9-98239h23h9-99h3ba993ba9h3ab","username": "someone","lease_duration": 14}`)}
				lease = NewLease(new(fakes.FakeCollection))
				lease.SetTask(new(taskmanager.Task))
				lease.InitFromHTTPRequest(request)
			})
			It("should update the task status", func() {
				controlStatus := lease.Task.Status
				lease.Procurement()
				Ω(lease.Task.Status).ShouldNot(Equal(controlStatus))
			})
		})

		Context("when calling with a valid lease (2c.small)", func() {
			var (
				lease   *Lease
				request *http.Request
			)
			BeforeEach(func() {
				request = new(http.Request)
				request.Body = fakes.FakeResponseBody{bytes.NewBufferString(`{"_id": "917397-292735-98293752935","sku":"2c.small", "inventory_id": "kaasd9sd9-98239h23h9-99h3ba993ba9h3ab","username": "someone","lease_duration": 14}`)}
				lease = NewLease(new(fakes.FakeCollection))
				lease.SetTask(new(taskmanager.Task))
				lease.InitFromHTTPRequest(request)
			})
			It("should update the task status", func() {
				controlStatus := lease.Task.Status
				lease.Procurement()
				Ω(lease.Task.Status).ShouldNot(Equal(controlStatus))
			})
		})

	})

	Describe(".ReStock()", func() {
		Context("when calling with a valid lease", func() {
			var (
				lease   *Lease
				request *http.Request
			)
			BeforeEach(func() {
				request = new(http.Request)
				request.Body = fakes.FakeResponseBody{bytes.NewBufferString(`{"_id": "917397-292735-98293752935","inventory_id": "kaasd9sd9-98239h23h9-99h3ba993ba9h3ab","username": "someone","lease_duration": 14}`)}
				lease = NewLease(new(fakes.FakeCollection))
				lease.SetTask(new(taskmanager.Task))
				lease.InitFromHTTPRequest(request)
			})
			It("should update the task status", func() {
				controlStatus := lease.Task.Status
				lease.ReStock()
				Ω(lease.Task.Status).ShouldNot(Equal(controlStatus))
			})
		})
	})

	Describe(".SetTask()", func() {
		Context("calling with a valid task on an initialized lease", func() {
			var (
				lease       = NewLease(new(fakes.FakeCollection))
				controlTask = new(taskmanager.Task)
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
