package heap

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestKLargest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test Suite for k-th largest")
}

var _ = Describe("The", func() {
	Describe("2-nd largest element in [1,2,3]", func() {
		arr := []int{1, 2, 3}
		k := 2
		It("should be 2", func() {
			Expect(kLargest(arr, k)).To(Equal(2))
		})
	})

	Describe("4-th largest element in [1,2,3]", func() {
		arr := []int{1, 2, 3}
		k := 4
		It("should return -1", func() {
			Expect(kLargest(arr, k)).To(Equal(-1))
		})
	})

	Describe("4-th largest element in [1,18,9,17,6]", func() {
		arr := []int{1, 18, 9, 17, 6}
		k := 4
		It("should be 6", func() {
			Expect(kLargest(arr, k)).To(Equal(6))
		})
	})
})
