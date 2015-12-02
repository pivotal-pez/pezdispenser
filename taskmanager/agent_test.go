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
	AgentTaskPollerInterval = time.Duration(0)
	Describe("given a NewAgent func", func() {
		Context("when the long running process exits without error", func() {

			var (
				controlAgent       *Agent
				controlTaskManager = &fakes.FakeTaskManager{
					ExpireEmitter: make(chan int64, 1),
				}
				controlCallerName = "fake-caller"
			)
			BeforeEach(func() {
				controlAgent = NewAgent(controlTaskManager, controlCallerName)
				controlTaskManager.SpyTaskSaved = new(Task)
				controlAgent.Run(func(*Agent) error {
					return nil
				})
			})
			It("then it should exit cleanly update status and expire the task", func() {

				select {
				case <-controlTaskManager.ExpireEmitter:
				default:
					Eventually(<-controlAgent.GetStatus()).Should(Equal(AgentTaskStatusRunning))
					Eventually(<-controlAgent.GetStatus()).Should(Equal(AgentTaskStatusComplete))
					close(controlTaskManager.ExpireEmitter)
				}
			})
		})
		Context("when the long running process exits with an error", func() {
			var (
				controlAgent       *Agent
				controlTaskManager = &fakes.FakeTaskManager{
					ExpireEmitter: make(chan int64, 1),
				}
				controlCallerName = "fake-caller"
			)
			BeforeEach(func() {
				controlAgent = NewAgent(controlTaskManager, controlCallerName)
				controlAgent.Run(func(*Agent) error {
					return errors.New("some random error")
				})
			})
			It("then it should exit w/ an error status", func() {
				select {
				case <-controlTaskManager.ExpireEmitter:
				default:
					Eventually(<-controlAgent.GetStatus()).Should(Equal(AgentTaskStatusRunning))
					Eventually(<-controlAgent.GetStatus()).Should(ContainSubstring(AgentTaskStatusFailed))
					close(controlTaskManager.ExpireEmitter)
				}
			})
		})

	})
	Describe("Given: a Run method", func() {
		Context("when called for a long running task", func() {
			var (
				controlAgent       *Agent
				controlTaskManager = &fakes.FakeTaskManager{
					ExpireEmitter: make(chan int64, 1),
				}
				controlCallerName = "fake-caller"
				controlExpires    = int64(10)
			)
			BeforeEach(func() {
				controlAgent = NewAgent(controlTaskManager, controlCallerName)
				controlAgent.GetTask().Expires = controlExpires
				controlAgent.Run(func(*Agent) error {
					time.Sleep(time.Duration(10) * time.Second)
					return nil
				})
			})
			It("then it should kick off a polling routine, which relays alive status at given interval offset", func() {
				lastCall := int64(controlExpires - 1)
				Consistently(func() bool {
					current := <-controlTaskManager.ExpireEmitter
					res := current >= lastCall
					lastCall = current
					return res
				}).Should(BeTrue())
			})
		})
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
