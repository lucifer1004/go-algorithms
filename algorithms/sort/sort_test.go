package sort

import (
	"math/rand"
	"sort"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type student struct {
	score int
	name  string
}

func TestSort(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test Suite for sort")
}

func compareInt(i, j interface{}) int {
	return i.(int) - j.(int)
}

func compareStudent(i, j interface{}) int {
	return i.(student).score - j.(student).score
}

var _ = Describe("Should sort", func() {
	It("fixed arrays", func() {
		h := array{-1, 2, 5, 3, 2, 4, 7}
		h.quickSort(compareInt)
		Expect(h).To(Equal(array{-1, 2, 2, 3, 4, 5, 7}))
	})

	Describe("random arrays", func() {
		const testCases = 10
		const maxNum = 100000
		for i := 0; i < testCases; i++ {
			It("should be sorted", func() {
				h := array{}
				num := rand.Intn(maxNum)
				for j := 0; j < num; j++ {
					h = append(h, rand.Int())
				}
				hClone := []int{}
				for j := range h {
					hClone = append(hClone, h[j].(int))
				}
				h.quickSort(compareInt)
				sort.Ints(hClone)
				for j := range h {
					Expect(h[j].(int)).To(Equal(hClone[j]))
				}
			})
		}
	})
})

var _ = Describe("Sort is stable on", func() {
	It("fixed arrays", func() {
		h := array{
			student{10, "Julia"},
			student{8, "Georgia"},
			student{8, "Gloria"},
			student{6, "Maria"},
		}
		h.quickSort(compareStudent)
		Expect(h[0].(student).score).To(Equal(6))
		Expect(h[2].(student).name).To(Equal("Gloria"))
	})
})
