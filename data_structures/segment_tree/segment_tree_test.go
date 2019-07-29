package segmenttree

import (
	"math/rand"

	"strconv"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSegmentTree(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test Suite for segment tree")
}

var _ = Describe("Segment tree should be correct", func() {
	const maxLen = 1e5
	const maxNum = 1e8
	for i := 1; i <= 20; i++ {
		It("Random test "+strconv.Itoa(i), func() {
			length := 100 + rand.Intn(maxLen)
			a := make([]int64, length)
			for j := range a {
				a[j] = int64(rand.Intn(maxNum))
			}
			s := newSegTree(a)
			for j := 0; j < 100; j++ {
				t := rand.Intn(10)
				l := int64(1 + rand.Intn(length/2))
				r := int64(length/2 + rand.Intn(length/2))
				if t%2 == 0 {
					sum := int64(0)
					for k := l - 1; k < r; k++ {
						sum += a[k]
					}
					Expect(s.query(l-1, r)).To(Equal(sum))
				} else {
					val := int64(rand.Intn(maxNum))
					for k := l - 1; k < r; k++ {
						a[k] += val
					}
					s.modify(l-1, r, val)
				}
			}
		})
	}
})
