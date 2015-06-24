package pezdispenser_test

import (
	"encoding/json"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/pivotal-pez/pezdispenser/service"

	"testing"
)

func TestOsutils(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Pez Dispenser Suite")
}

func isValidResponseMessage(res, control string) {
	msg := new(ResponseMessage)
	json.Unmarshal([]byte(res), msg)
	Ω(msg.Body).ShouldNot(BeNil())
	Ω(msg.Status).ShouldNot(BeNil())
	Ω(msg.Version).ShouldNot(BeNil())
	Ω(msg.Body).Should(ContainSubstring(control))
}
