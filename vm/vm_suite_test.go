package vm_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestVendingMachine(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Vending Machine Suite")
}
