package bsearch

// Search for integer `val` in an ascending and non-duplicate array `arr`
// If `val` exists, return its index, otherwise return -1
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

// Search for the first integer larger than `val` in an ascending array `arr`
// If such integer exists, return its value, otherwise return -1
func binarySearchFirstLargerNumber(arr []int, val int) int {
	l := 0
	r := len(arr) - 1
	for l <= r {
		mid := l + (r-l)/2
		if arr[mid] > val {
			if mid == 0 || arr[mid-1] <= val {
				return arr[mid]
			}
			r = mid + 1
		} else if arr[mid] <= val {
			l = mid + 1
		}
	}
	return -1
}

// Search for integer `val` in an ascending array `arr`
// If `val` exists, return the first and last indices, otherwise return [-1, -1]
func binarySearchFirstAndLast(arr []int, val int) [2]int {
	first := -1
	last := -1

	l := 0
	r := len(arr) - 1
	for l <= r {
		mid := l + (r-l)/2
		if arr[mid] == val {
			if mid == 0 || arr[mid-1] < val {
				first = mid
			}
			r = mid - 1
		} else if arr[mid] < val {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	l = 0
	r = len(arr) - 1
	for l <= r {
		mid := l + (r-l)/2
		if arr[mid] == val {
			if mid == len(arr)-1 || arr[mid+1] > val {
				last = mid
			}
			l = mid + 1
		} else if arr[mid] < val {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return [2]int{first, last}
}
