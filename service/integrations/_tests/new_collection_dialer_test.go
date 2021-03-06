package integrations_test

import (
	"fmt"
	"os"
	"time"

	"labix.org/v2/mgo/bson"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pivotal-pez/pezdispenser/fakes"
	. "github.com/pivotal-pez/pezdispenser/service/integrations"
	"github.com/pivotal-pez/pezdispenser/taskmanager"
)

var _ = Describe("GetTaskByIdController()", func() {
	Describe("NewCollectionDialer", func() {
		Context("when the collection dialer is pointing at a real mongo integration point", func() {
			var (
				col            Collection
				err            error
				ip             = os.Getenv("MONGO_PORT_27017_TCP_ADDR")
				port           = os.Getenv("MONGO_PORT_27017_TCP_PORT")
				databaseName   string
				collectionName string
				mongoURI       string
			)

			BeforeEach(func() {
				databaseName = bson.NewObjectId().Hex()
				collectionName = bson.NewObjectId().Hex()
				mongoURI = fmt.Sprintf("mongodb://%s:%s/%s", ip, port, databaseName)
				col, err = NewCollectionDialer(mongoURI, databaseName, collectionName)
			})

			AfterEach(func() {
				col.Close()
			})

			It("should be able to connect", func() {
				Ω(err).ShouldNot(HaveOccurred())
			})

			Describe(".FindAndModify()", func() {
				controlCallerName := "fake.caller"
				controlStatus := "doing---stuff"
				controlEndStatus := "closed"
				Context("when the given selector finds a match", func() {
					var (
						controlTask *taskmanager.Task
						tm          *taskmanager.TaskManager
					)
					BeforeEach(func() {
						tm = taskmanager.NewTaskManager(col)
						controlTask = tm.NewTask(controlCallerName, taskmanager.TaskLongPollQueue, controlStatus)
						controlTask.Expires = time.Now().UnixNano()
						tm.SaveTask(controlTask)
					})
					AfterEach(func() {
						controlTask.Status = controlEndStatus
						tm.SaveTask(controlTask)
					})
					It("should update the object and return its value", func() {
						nowEpoch := time.Now().UnixNano()
						taskOutput := new(taskmanager.Task)
						ci, err := col.FindAndModify(
							bson.M{
								"caller_name": controlCallerName,
								"expires":     bson.M{"$lte": nowEpoch},
							},
							bson.M{
								"expires": time.Now().Add(5 * time.Minute).UnixNano(),
							},
							taskOutput,
						)
						Ω(err).ShouldNot(HaveOccurred())
						Ω(ci.Updated).Should(Equal(1))
						var equal = controlTask.Equal(*taskOutput)
						Ω(equal).Should(Equal(true))
					})
				})
			})

			Describe("UpsertID()", func() {
				Context("called with a valid ID and update object", func() {
					It("should be able to upsert a record", func() {
						controlID := bson.NewObjectId()
						upsertdata := &fakes.FakeTask{
							ID:     controlID,
							Status: "fake status",
						}
						controlCount, _ := col.Count()
						info, err := col.UpsertID(controlID, upsertdata)
						count, _ := col.Count()
						Ω(err).ShouldNot(HaveOccurred())
						Ω(info.UpsertedId.(bson.ObjectId).Hex()).Should(Equal(controlID.Hex()))
						Ω(count).ShouldNot(Equal(controlCount))
					})
				})
			})

			Describe("FindOne()", func() {
				Context("when called with a valid formatted string ID and a result object pointer", func() {
					It("should be able read a given ID and apply the values to the result object", func() {
						controlID := bson.NewObjectId()
						upsertdata := &fakes.FakeTask{
							ID:     controlID,
							Status: "fake status",
						}
						resultData := new(fakes.FakeTask)
						col.UpsertID(controlID, upsertdata)
						err := col.FindOne(controlID.Hex(), resultData)
						Ω(err).ShouldNot(HaveOccurred())
						Ω(resultData.Status).Should(Equal(upsertdata.Status))
						Ω(resultData.ID).Should(Equal(upsertdata.ID))
						Ω(resultData.Timestamp).Should(Equal(upsertdata.Timestamp))
					})
				})
				Context("when called with a an invalid ID format", func() {
					It("should return an invalid id format error", func() {
						controlID := "badformat"
						resultData := new(fakes.FakeTask)
						err := col.FindOne(controlID, resultData)
						Ω(err).Should(HaveOccurred())
						Ω(err).Should(Equal(ErrInvalidID))
					})
				})
			})
		})
	})
})
