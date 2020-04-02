package controller_test

import (
	"fmt"
	sriovTest "github.com/fafucoder/sriov-crd/testing"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestController(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Controller Suite")
}

var _ = BeforeSuite(func() {
	err := sriovTest.CreateTmpSysFs()

	if err != nil {
		fmt.Errorf("error is, %v", err)
		panic(err)
	}
})

var _ = AfterSuite(func() {
	err := sriovTest.RemoveTmpSysFs()
	if err != nil {
		fmt.Errorf("error is, %v", err)
		panic(err)
	}
})
