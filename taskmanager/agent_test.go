package taskmanager_test

import (
	. "github.com/pivotal-pez/pezdispenser/taskmanager"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Agent", func() {
	Describe("given a NewAgent func", func() {
		Context("when called with a given a task", func() {
			var (
				controlAgent *Agent
				controlTask  = new(TaskManager).NewTask("", "", "")
			)
			BeforeEach(func() {
				controlAgent = NewAgent(controlTask)
			})
			It("then it should leverage a pre-initialized task passed by the caller and it should immediately return the task object	and not block", func() {
				Ω(controlAgent.Run(func(*Agent) {})).Should(Equal(controlTask))
			})
		})
	})
	Describe("Given: a Run method", func() {
		Context("when called", func() {
			It("then it should inject the agent context and call the given function", func() {
				agentSpy := make(chan *Agent)
				a := new(Agent)
				a.Run(func(localAgent *Agent) {
					agentSpy <- localAgent
				})
				Eventually(<-agentSpy).Should(Equal(a))
			})
		})
	})
})
