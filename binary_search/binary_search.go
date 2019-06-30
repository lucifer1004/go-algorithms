package algo

func binarySearch(arr []int, val int) int {
	l := 0
	r := len(arr) - 1
	for l <= r {
		mid := l + (r-l)/2
		if arr[mid] > val {
			r = mid - 1
		} else if arr[mid] < val {
			l = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
