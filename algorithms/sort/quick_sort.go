package sort

import "math/rand"

func (a array) quickSort(cmp compareFunc) {
	if len(a) < 2 {
		return
	}

	left, right := 0, len(a)-1

	pivot := rand.Int() % len(a)

	a[pivot], a[right] = a[right], a[pivot]

	for i := range a {
		if cmp(a[i], a[right]) < 0 || (cmp(a[i], a[right]) == 0 && i < pivot) {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	a[:left].quickSort(cmp)
	a[left+1:].quickSort(cmp)
}
