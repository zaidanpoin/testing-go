package kalkulator_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestKalkulator(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Kalkulator Suite")
}
