package taskmanager_test

import (
	"time"

	"labix.org/v2/mgo"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pivotal-pez/pezdispenser/fakes"
	. "github.com/pivotal-pez/pezdispenser/taskmanager"
	"labix.org/v2/mgo/bson"
)

var _ = Describe("TaskManager", func() {
	Describe("Given: .GarbageCollectExpiredAgents method", func() {

		var tm *TaskManager

		Context("when: called and it finds expired agents", func() {
			var controlUpdateCount = 25
			BeforeEach(func() {
				tm = NewTaskManager(&fakes.FakeCollection{
					FakeChangeInfo: &mgo.ChangeInfo{
						Updated: controlUpdateCount,
					},
				})
			})
			It("then: it should expire them", func() {
				gcList, err := tm.GarbageCollectExpiredAgents("")
				Ω(err).ShouldNot(HaveOccurred())
				Ω(gcList.Updated).ShouldNot(BeNil())
			})
		})
	})
	Describe("Given: .FindAndStallTaskForCaller()", func() {
		var tm *TaskManager

		BeforeEach(func() {
			tm = NewTaskManager(new(fakes.FakeCollection))
		})
		Context("When: call yields no results", func() {
			It("Then: it should return a no-results error", func() {
				_, err := tm.FindAndStallTaskForCaller("")
				Ω(err).Should(HaveOccurred())
				Ω(err).Should(Equal(ErrNoResults))
			})
		})
	})

	Describe(".FindTask()", func() {
		var tm *TaskManager

		BeforeEach(func() {
			tm = NewTaskManager(new(fakes.FakeCollection))
		})
		Context("when called", func() {
			It("Should do nothing right now", func() {
				tm.FindTask("")
				Ω(true).Should(BeTrue())
			})
		})
	})

	Describe(".NewTask()", func() {
		var tm *TaskManager

		BeforeEach(func() {
			tm = NewTaskManager(new(fakes.FakeCollection))
		})
		Context("when called", func() {
			It("Should return a newly initialized task", func() {
				controlTimestampEpoch := time.Now().UnixNano()
				t := tm.NewTask("random.skutype", TaskLongPollQueue, "processing")
				Ω(t.ID.Hex()).ShouldNot(BeEmpty())
				Ω(t.Timestamp).Should(BeNumerically(">", controlTimestampEpoch))
			})
		})
	})

	Describe(".SaveTask()", func() {
		var tm *TaskManager

		BeforeEach(func() {
			tm = NewTaskManager(new(fakes.FakeCollection))
		})
		Context("when given an existing task", func() {
			It("should update the task", func() {
				controlID := bson.NewObjectId()
				task := &Task{
					ID: controlID,
				}
				t, err := tm.SaveTask(task)
				Ω(err).ShouldNot(HaveOccurred())
				Ω(t.ID).Should(Equal(controlID))
			})
		})

		Context("when given a new task", func() {
			It("should create an id and save it", func() {
				task := new(Task)
				controlID := task.ID
				Ω(task.ID.Hex()).Should(BeEmpty())
				t, err := tm.SaveTask(task)
				Ω(err).ShouldNot(HaveOccurred())
				Ω(task.ID.Hex()).ShouldNot(BeEmpty())
				Ω(t.ID).ShouldNot(Equal(controlID))
			})
		})
	})
})
