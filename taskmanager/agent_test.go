package taskmanager_test

import (
	. "github.com/pivotal-pez/pezdispenser/taskmanager"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Agent", func() {
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
