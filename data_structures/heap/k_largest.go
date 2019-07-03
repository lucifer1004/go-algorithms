package heap

func kLargest(arr []int, k int) int {
	if k > len(arr) || k < 1 {
		return -1
	}

	h := buildHeap(arr[0:k])

	for i := k; i < len(arr); i++ {
		if arr[i] > queryMin(h) {
			deleteMin(h)
			insert(h, arr[i])
		}
	}

	return queryMin(h)
}
