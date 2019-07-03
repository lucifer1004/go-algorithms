package fastmul

import (
	"math/rand"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const modulo = 1000000007

func TestFastEXP(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test Suite for fast mul")
}

var _ = Describe("Calculate", func() {
	Describe("Multiplication of", func() {
		for i := 0; i < 10; i++ {
			It("Random numbers", func() {
				x := rand.Intn(modulo)
				y := rand.Intn(modulo)
				a1 := fastMul(x, y, modulo)
				a2 := x * y % modulo
				Expect(a1).To(Equal(a2))
			})
		}
	})
})
