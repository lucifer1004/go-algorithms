package heap

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestHeap(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test Suite for heap")
}

var _ = Describe("Heap should be correct", func() {
	var h *intHeap
	arr := []int{1, 4, 7, 9}

	BeforeEach(func() {
		h = BuildHeap(arr)
	})

	Describe("After", func() {
		It("Each operation", func() {
			Insert(h, 0)
			Expect(DeleteMin(h)).To(Equal(0))
			Expect(DeleteMin(h)).To(Equal(1))
			Insert(h, 5)
			Insert(h, 9)
			Expect(DeleteMin(h)).To(Equal(4))
			Expect(DeleteMin(h)).To(Equal(5))
			Expect(DeleteMin(h)).To(Equal(7))
			Expect(DeleteMin(h)).To(Equal(9))
			Expect(DeleteMin(h)).To(Equal(9))
			Expect(DeleteMin(h)).To(Equal(-1))
		})
	})
})
