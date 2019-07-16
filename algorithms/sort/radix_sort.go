package sort

func (a uintArray) countingSort(b0, m uint) {
	n := len(a)
	b := make(uintArray, n)
	nb := m + 1
	m <<= b0
	cv := make([]uint, nb)
	for i := 0; i < n; i++ {
		x := (a[i] & m) >> b0
		cv[x]++
	}
	for i := 1; i < int(nb); i++ {
		cv[i] += cv[i-1]
	}
	k := n - 1
	for k >= 0 {
		x := (a[k] & m) >> b0
		cv[x]--
		i := cv[x]
		b[i] = a[k]
		k--
	}
	for i := range a {
		a[i] = b[i]
	}
}

func (a uintArray) radixSort() {
	var nb uint = 8
	var tnb uint = 32
	var m uint = 1<<nb - 1
	for k, b0 := 1, uint(0); b0 < tnb; k, b0 = k+1, b0+nb {
		a.countingSort(b0, m)
	}
}
