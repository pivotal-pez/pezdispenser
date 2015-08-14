package vcloud_client_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestVCDClient(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "VCD Client Suite")
}
