package cornfields

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCornFields(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test Suite for Corn Fields")
}

var _ = Describe("The ways of placing cows modulo 100000000", func() {
	var a = [][]int{[]int{0}}

	var b = [][]int{
		[]int{1, 0, 0, 0, 0},
		[]int{1, 0, 1, 0, 1},
		[]int{1, 0, 0, 0, 1},
	}

	var c = [][]int{
		[]int{1, 1, 1, 1},
		[]int{1, 1, 1, 1},
		[]int{1, 1, 1, 1},
		[]int{1, 1, 1, 1},
	}

	var d = [][]int{
		[]int{1, 1, 1},
		[]int{0, 1, 0},
	}

	Describe("for farmland A should", func() {
		It("equal to 1", func() {
			Expect(countWays(a)).To(Equal(1))
		})
	})

	Describe("for farmland B should", func() {
		It("equal to 30", func() {
			Expect(countWays(b)).To(Equal(30))
		})
	})

	Describe("for farmland C should", func() {
		It("equal to 1234", func() {
			Expect(countWays(c)).To(Equal(1234))
		})
	})

	Describe("for farmland D should", func() {
		It("equal to 9", func() {
			Expect(countWays(d)).To(Equal(9))
		})
	})
})
