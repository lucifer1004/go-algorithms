package heap

func kLargest(arr []int, k int) int {
	if k > len(arr) || k < 1 {
		return -1
	}

	h := BuildHeap(arr[0:k])

	for i := k; i < len(arr); i++ {
		if arr[i] > QueryMin(h) {
			DeleteMin(h)
			Insert(h, arr[i])
		}
	}

	return QueryMin(h)
}
