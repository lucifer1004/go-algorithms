package algo

import (
	"math/rand"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMinWindow(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test Suite for binary search")
}

var _ = Describe("Binary search", func() {
	Describe("2 in [1,2,3]", func() {
		arr := []int{1, 2, 3}
		val := 2
		It("should return 1", func() {
			Expect(binarySearch(arr, val)).To(Equal(1))
		})
	})

	Describe("0 in [1,2,3]", func() {
		arr := []int{1, 2, 3}
		val := 0
		It("should return -1", func() {
			Expect(binarySearch(arr, val)).To(Equal(-1))
		})
	})

	Describe("random number in random array", func() {
		for i := 0; i < 10; i++ {
			arrLen := rand.Intn(1000)
			arr := make([]int, arrLen)
			for j := 0; j < arrLen; j++ {
				arr[j] = j
			}
			val := rand.Intn(arrLen)
			It("should return val", func() {
				Expect(binarySearch(arr, val)).To(Equal(val))
			})
		}
	})
})
