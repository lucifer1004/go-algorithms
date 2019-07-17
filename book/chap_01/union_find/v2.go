package unionfind

// IntUnionFindV2 ...
type IntUnionFindV2 struct {
	id []int
}

func (intUF *IntUnionFindV2) init(n int) {
	(*intUF).id = make([]int, n)
	for i := 0; i < n; i++ {
		(*intUF).id[i] = i
	}
}

func (intUF *IntUnionFindV2) union(p, q int) {
	pRoot := intUF.find(p)
	qRoot := intUF.find(q)
	(*intUF).id[pRoot] = qRoot
}

func (intUF *IntUnionFindV2) find(p int) int {
	for (*intUF).id[p] != p {
		p = (*intUF).id[p]
	}
	return p
}

func (intUF *IntUnionFindV2) isConnected(p, q int) bool {
	return intUF.find(p) == intUF.find(q)
}

func (intUF *IntUnionFindV2) count() int {
	total := 0
	for i, v := range (*intUF).id {
		if i == v {
			total++
		}
	}
	return total
}
