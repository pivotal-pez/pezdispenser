package taskmanager_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pivotal-pez/pezdispenser/fakes"
	. "github.com/pivotal-pez/pezdispenser/taskmanager"
	"labix.org/v2/mgo/bson"
)

var _ = Describe("TaskManager", func() {
	Describe(".FindLockFirstCallerName()", func() {
		var tm *TaskManager

		BeforeEach(func() {
			tm = NewTaskManager(new(fakes.FakeCollection))
		})
		Context("when called", func() {
			It("Should do nothing right now", func() {
				tm.FindLockFirstCallerName("")
				Ω(true).Should(BeTrue())
			})
		})
	})

	Describe(".UnLockTask()", func() {
		var tm *TaskManager

		BeforeEach(func() {
			tm = NewTaskManager(new(fakes.FakeCollection))
		})
		Context("when called", func() {
			It("Should do nothing right now", func() {
				tm.UnLockTask("")
				Ω(true).Should(BeTrue())
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
				t := tm.NewTask("random.skutype", TaskLongPollQueue, "processing")
				Ω(t.ID.Hex()).ShouldNot(BeEmpty())
				Ω(t.Timestamp).ShouldNot(Equal(time.Time{}))
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
