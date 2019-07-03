package heap

import (
	"container/heap"
)

type intHeap []int

func (h intHeap) Len() int {
	return len(h)
}

func (h intHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h intHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *intHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func insert(h *intHeap, val int) {
	heap.Push(h, val)
}

func buildHeap(arr []int) *intHeap {
	h := &intHeap{}
	heap.Init(h)
	for _, v := range arr {
		insert(h, v)
	}
	return h
}

func queryMin(h *intHeap) int {
	if h.Len() > 0 {
		return (*h)[0]
	}
	return -1
}

func deleteMin(h *intHeap) int {
	if h.Len() > 0 {
		return heap.Pop(h).(int)
	}
	return -1
}
