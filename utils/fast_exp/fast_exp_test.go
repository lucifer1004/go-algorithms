package fastexp

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const modulo = 1000000007

func TestFastEXP(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test Suite for fast exp")
}

var _ = Describe("", func() {
	Describe("2^8", func() {
		It("should be 256", func() {
			Expect(fastExp(2, 8, modulo)).To(Equal(256))
		})
	})
})
