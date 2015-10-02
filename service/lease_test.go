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
	"github.com/pivotal-pez/pezdispenser/skus"
	"github.com/pivotal-pez/pezdispenser/taskmanager"
)

var _ = Describe("Lease", func() {
	Describe("NewLease", func() {
		Context("calling new lease", func() {
			It("should yield a lease valid lease", func() {
				lease := NewLease(new(fakes.FakeCollection), map[string]skus.Sku{
					"": new(fakes.FakeSku),
				})
				Ω(lease).ShouldNot(BeNil())
			})
		})
	})

	Describe(".InventoryAvailable()", func() {
		Context("When: no task matches can be found yielding a nil changeInfo", func() {
			var (
				lease *Lease
			)

			BeforeEach(func() {
				lease = NewLease(fakes.NewFakeCollection(fakes.FakeCollectionHasNilChangeInfo), map[string]skus.Sku{
					"": new(fakes.FakeSku),
				})
			})

			It("Then: it should not panic", func() {
				Ω(func() {
					lease.InventoryAvailable()
				}).ShouldNot(Panic())
			})
		})

		Context("when we can find the record but taskStatus is != available", func() {
			var (
				lease *Lease
			)
			BeforeEach(func() {
				lease = NewLease(fakes.NewFakeCollection(fakes.FakeCollectionHasNoChanges), map[string]skus.Sku{
					"": new(fakes.FakeSku),
				})
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
				col := fakes.NewFakeCollection(fakes.FakeCollectionHasChanges)
				col.ControlTask = taskmanager.Task{
					Status: TaskStatusAvailable,
				}
				lease = NewLease(col, map[string]skus.Sku{
					"": new(fakes.FakeSku),
				})
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
				col := fakes.NewFakeCollection(fakes.FakeCollectionHasNilChangeInfo)
				col.ErrFindAndModify = mgo.ErrNotFound
				col.ErrControl = mgo.ErrNotFound
				lease = NewLease(col, map[string]skus.Sku{
					"": new(fakes.FakeSku),
				})
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
				lease = NewLease(fakes.NewFakeCollection(fakes.FakeCollectionHasChanges), map[string]skus.Sku{
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
				lease = NewLease(fakes.NewFakeCollection(fakes.FakeCollectionHasChanges), map[string]skus.Sku{
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
				lease = NewLease(fakes.NewFakeCollection(fakes.FakeCollectionHasChanges), map[string]skus.Sku{
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
				lease = NewLease(fakes.NewFakeCollection(fakes.FakeCollectionHasChanges), map[string]skus.Sku{
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

	Describe("Given: method .InventoryAvailable()", func() {
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
				lease = NewLease(fakes.NewFakeCollection(fakes.FakeCollectionHasChanges), map[string]skus.Sku{
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
				lease = NewLease(fakes.NewFakeCollection(fakes.FakeCollectionHasChanges), map[string]skus.Sku{
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
