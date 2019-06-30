package bsearch

import (
	"math/rand"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBinarySearch(t *testing.T) {
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

var _ = Describe("Binary search first number larger than", func() {
	Describe("2 in [1,2,3]", func() {
		arr := []int{1, 2, 3}
		val := 2
		It("should return 3", func() {
			Expect(binarySearchFirstLargerNumber(arr, val)).To(Equal(3))
		})
	})

	Describe("3 in [1,2,3]", func() {
		arr := []int{1, 2, 3}
		val := 3
		It("should return -1", func() {
			Expect(binarySearchFirstLargerNumber(arr, val)).To(Equal(-1))
		})
	})
})

var _ = Describe("Binary search first and last occurrence of", func() {
	Describe("2 in [1,2,2,2,3]", func() {
		arr := []int{1, 2, 2, 2, 3}
		val := 2
		It("should return [1, 3]", func() {
			Expect(binarySearchFirstAndLast(arr, val)).To(Equal([2]int{1, 3}))
		})
	})

	Describe("3 in [1,2,3]", func() {
		arr := []int{1, 2, 3}
		val := 3
		It("should return [2, 2]", func() {
			Expect(binarySearchFirstAndLast(arr, val)).To(Equal([2]int{2, 2}))
		})
	})

	Describe("0 in [1,2,3]", func() {
		arr := []int{1, 2, 3}
		val := 0
		It("should return [-1, -1]", func() {
			Expect(binarySearchFirstAndLast(arr, val)).To(Equal([2]int{-1, -1}))
		})
	})
})
