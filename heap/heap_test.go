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
		h = buildHeap(arr)
	})

	Describe("After", func() {
		It("Each operation", func() {
			insert(h, 0)
			Expect(deleteMin(h)).To(Equal(0))
			Expect(deleteMin(h)).To(Equal(1))
			insert(h, 5)
			insert(h, 9)
			Expect(deleteMin(h)).To(Equal(4))
			Expect(deleteMin(h)).To(Equal(5))
			Expect(deleteMin(h)).To(Equal(7))
			Expect(deleteMin(h)).To(Equal(9))
			Expect(deleteMin(h)).To(Equal(9))
			Expect(deleteMin(h)).To(Equal(-1))
		})
	})
})
