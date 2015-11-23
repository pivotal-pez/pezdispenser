package taskmanager_test

import (
	"errors"
	"time"

	"github.com/pivotal-pez/pezdispenser/fakes"
	. "github.com/pivotal-pez/pezdispenser/taskmanager"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Agent", func() {
	Describe("given a NewAgent func", func() {
		Context("when called with a given a task", func() {

			var (
				controlAgent       *Agent
				controlTaskManager = new(fakes.FakeTaskManager)
				controlCallerName  = "fake-caller"
			)
			BeforeEach(func() {
				controlAgent = NewAgent(controlTaskManager, controlCallerName)
			})
			It("then it should leverage a pre-initialized task passed by the caller", func() {
				controlAgent.Run(func(*Agent) (err error) { return })
				Ω(controlAgent.GetTask()).ShouldNot(BeNil())
				Ω(controlAgent.GetTask().CallerName).Should(Equal(controlCallerName))
			})

			It("then it should not block, executing the function in the background", func() {
				controlAgent.Run(func(*Agent) error {
					time.Sleep(time.Duration(10) * time.Second)
					return nil
				})
				Ω(controlAgent.GetTask().Status).Should(Equal(AgentTaskStatusRunning))
			})
		})
		Context("when the long running process exits without error", func() {

			var (
				controlAgent       *Agent
				controlTaskManager = new(fakes.FakeTaskManager)
				controlCallerName  = "fake-caller"
			)
			BeforeEach(func() {
				controlAgent = NewAgent(controlTaskManager, controlCallerName)
				controlTaskManager.SpyTaskSaved = new(Task)
			})
			It("then it should exit cleanly update status and expire the task", func() {
				controlAgent.Run(func(*Agent) error {
					return nil
				})
				Eventually(<-controlAgent.GetStatus()).Should(Equal(AgentTaskStatusRunning))
				Eventually(<-controlAgent.GetStatus()).Should(Equal(AgentTaskStatusComplete))
				Ω(controlTaskManager.SpyTaskSaved.Expires).Should(Equal(int64(0)))
			})
		})
		Context("when the long running process exits with an error", func() {

			var (
				controlAgent       *Agent
				controlTaskManager = new(fakes.FakeTaskManager)
				controlCallerName  = "fake-caller"
			)
			BeforeEach(func() {
				controlAgent = NewAgent(controlTaskManager, controlCallerName)
				controlTaskManager.SpyTaskSaved = new(Task)
			})
			It("then it should exit w/ an error status", func() {
				controlAgent.Run(func(*Agent) error {
					return errors.New("some random error")
				})
				Eventually(<-controlAgent.GetStatus()).Should(Equal(AgentTaskStatusRunning))
				Eventually(<-controlAgent.GetStatus()).Should(ContainSubstring(AgentTaskStatusFailed))
				Ω(controlTaskManager.SpyTaskSaved.Expires).Should(Equal(int64(0)))
			})
		})

	})
	Describe("Given: a Run method", func() {
		Context("when called", func() {
			It("then it should inject the agent context and call the given function", func() {
				agentSpy := make(chan *Agent)
				a := NewAgent(new(fakes.FakeTaskManager), "fake-caller")
				a.Run(func(localAgent *Agent) error {
					agentSpy <- localAgent
					return nil
				})
				Eventually(<-agentSpy).Should(Equal(a))
			})
		})
	})
})
