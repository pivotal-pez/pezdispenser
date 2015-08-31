package taskmanager_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/pivotal-pez/pezdispenser/taskmanager"
)

var _ = Describe("Task", func() {
	Describe(".GetRedactedVersion()", func() {
		Context("when called on a task", func() {
			var (
				controlPrivateMetaKey = "random"
				controlPublicMetaKey  = "hithere"
				controlCaller         = "caller_nnnn"
				controlStatus         = "somestatus"
				task                  = NewTaskManager(nil).NewTask(controlCaller, TaskLongPollQueue, controlStatus)
			)
			BeforeEach(func() {
				task.SetPrivateMeta(controlPrivateMetaKey, "random-private")
				task.SetPublicMeta(controlPublicMetaKey, "random-public")
			})
			It("should return a RedactedTask, containing task info, excluding private info", func() {
				redactedTask := task.GetRedactedVersion()
				Ω(redactedTask).ShouldNot(Equal(*task))
			})
		})
	})
})
