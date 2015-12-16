package pdclient_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestPDClient(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Pez Dispenser Client")
}
