package utils_test

import (
	sriovTest "github.com/fafucoder/sriov-crd/testing"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestUtils(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Utils Suite")
}

var _ = BeforeSuite(func() {
	err := sriovTest.CreateTmpSysFs()
	if err != nil {
		panic(err)
	}
})

var _ = AfterSuite(func() {
	err := sriovTest.RemoveTmpSysFs()
	if err != nil {
		panic(err)
	}
})
