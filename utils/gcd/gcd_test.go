package utils

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGCD(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test Suite for GCD")
}

var _ = Describe("GCD of", func() {
	It("(19, 8) is 1", func() {
		Expect(gcd(19, 8)).To(Equal(1))
	})

	It("(182, 2) is 2", func() {
		Expect(gcd(182, 2)).To(Equal(2))
	})
})
