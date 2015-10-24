package pezdispenser_test

import (
	"bytes"
	"net/http"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pivotal-pez/pezdispenser/fakes"
	. "github.com/pivotal-pez/pezdispenser/service"
	"github.com/pivotal-pez/pezdispenser/skurepo"
	"github.com/pivotal-pez/pezdispenser/taskmanager"
)

var _ = Describe("Lease", func() {
	Describe("NewLease", func() {
		Context("calling new lease", func() {
			It("should yield a lease valid lease", func() {
				lease := NewLease(new(fakes.FakeCollection), map[string]skurepo.Sku{
					"": new(fakes.FakeSku),
				})
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
				lease = NewLease(fakes.NewFakeCollection(fakes.FakeCollectionHasChanges), map[string]skurepo.Sku{
					"": new(fakes.FakeSku),
				})
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
				request.Body = fakes.FakeResponseBody{bytes.NewBufferString(`{"sku":"sku1234", "lease_id": "917397-292735-98293752935","inventory_id": "kaasd9sd9-98239h23h9-99h3ba993ba9h3ab","username": "someone","lease_duration": 14}`)}
				lease = NewLease(fakes.NewFakeCollection(fakes.FakeCollectionHasChanges), map[string]skurepo.Sku{
					"sku1234": &fakes.FakeSku{
						ProcurementTask: &taskmanager.Task{
							Status: TaskStatusUnavailable,
						},
					},
				})
				lease.LeaseEndDate = time.Now().UnixNano()
				lease.ProcurementMeta = make(map[string]interface{})
			})

			It("should return the lease object as the response", func() {
				statusCode, response := lease.Post(fakes.MockLogger, request)
				Ω(statusCode).Should(Equal(http.StatusCreated))
				Ω(response.(*Lease).Task.Status).Should(Equal(TaskStatusUnavailable))
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
				request.Body = fakes.FakeResponseBody{bytes.NewBufferString(`{"lease_id": "917397-292735-98293752935","inventory_id": "kaasd9sd9-98239h23h9-99h3ba993ba9h3ab","username": "someone","lease_duration": 14}`)}
				lease = NewLease(fakes.NewFakeCollection(fakes.FakeCollectionHasChanges), map[string]skurepo.Sku{
					"": new(fakes.FakeSku),
				})
			})

			It("should populate the lease's fields with the given lease", func() {
				err := lease.InitFromHTTPRequest(request)
				Ω(err).ShouldNot(HaveOccurred())
				Ω(lease.ID).Should(Equal("917397-292735-98293752935"))
				Ω(lease.InventoryID).Should(Equal("kaasd9sd9-98239h23h9-99h3ba993ba9h3ab"))
				Ω(lease.UserName).Should(Equal("someone"))
				Ω(lease.LeaseDuration).Should(Equal(float64(14)))
				Ω(lease.Task).ShouldNot(BeNil())
			})
		})
		Context("called with an invalid request body", func() {
			var (
				lease   *Lease
				request *http.Request
			)

			BeforeEach(func() {
				lease = NewLease(fakes.NewFakeCollection(fakes.FakeCollectionHasChanges), map[string]skurepo.Sku{
					"": new(fakes.FakeSku),
				})
				request = new(http.Request)
			})

			It("should return a error", func() {
				err := lease.InitFromHTTPRequest(request)
				Ω(err).Should(HaveOccurred())
			})
		})
	})

	Describe(".Procurement()", func() {
		Context("when calling with a valid lease", func() {
			var (
				lease   *Lease
				request *http.Request
			)
			BeforeEach(func() {
				request = new(http.Request)
				request.Body = fakes.FakeResponseBody{bytes.NewBufferString(`{"sku":"1234","lease_id": "917397-292735-98293752935","sku":"2c.small", "inventory_id": "kaasd9sd9-98239h23h9-99h3ba993ba9h3ab","username": "someone","lease_duration": 14}`)}
				lease = NewLease(fakes.NewFakeCollection(fakes.FakeCollectionHasChanges), map[string]skurepo.Sku{
					"1234": &fakes.FakeSku{
						ProcurementTask: &taskmanager.Task{
							Status: TaskStatusRestocking,
						},
					},
				})
				lease.LeaseEndDate = time.Now().UnixNano()
				lease.ProcurementMeta = make(map[string]interface{})
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
				request.Body = fakes.FakeResponseBody{bytes.NewBufferString(`{"sku":"1234","lease_id": "917397-292735-98293752935","inventory_id": "kaasd9sd9-98239h23h9-99h3ba993ba9h3ab","username": "someone","lease_duration": 14}`)}
				lease = NewLease(fakes.NewFakeCollection(fakes.FakeCollectionHasChanges), map[string]skurepo.Sku{
					"1234": &fakes.FakeSku{
						ReStockTask: &taskmanager.Task{
							Status: TaskStatusRestocking,
						},
					},
				})
				lease.InitFromHTTPRequest(request)
			})
			It("should update the task status", func() {
				lease.ReStock()
				Ω(lease.Task.Status).Should(Equal(TaskStatusRestocking))
			})
		})
	})
})
