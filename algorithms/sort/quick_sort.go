package sort

import "math/rand"

type array []interface{}

func (a array) quickSort(compareFunc func(i, j interface{}) int) {
	if len(a) < 2 {
		return
	}

	left, right := 0, len(a)-1

	pivot := rand.Int() % len(a)

	a[pivot], a[right] = a[right], a[pivot]

	for i := range a {
		if compareFunc(a[i], a[right]) < 0 || (compareFunc(a[i], a[right]) == 0 && i < pivot) {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	a[:left].quickSort(compareFunc)
	a[left+1:].quickSort(compareFunc)
}
