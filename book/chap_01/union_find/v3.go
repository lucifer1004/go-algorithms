package unionfind

// IntUnionFindV3 ...
type IntUnionFindV3 struct {
	id   []int
	size []int
}

func (intUF *IntUnionFindV3) init(n int) {
	(*intUF).id = make([]int, n)
	(*intUF).size = make([]int, n)
	for i := 0; i < n; i++ {
		(*intUF).id[i] = i
		(*intUF).size[i] = 1
	}
}

func (intUF *IntUnionFindV3) union(p, q int) {
	pRoot := intUF.find(p)
	qRoot := intUF.find(q)
	if (*intUF).size[pRoot] < (*intUF).size[qRoot] {
		(*intUF).id[pRoot] = qRoot
		(*intUF).size[qRoot] += (*intUF).size[pRoot]
	} else {
		(*intUF).id[qRoot] = pRoot
		(*intUF).size[pRoot] += (*intUF).size[qRoot]
	}
}

func (intUF *IntUnionFindV3) find(p int) int {
	for (*intUF).id[p] != p {
		p = (*intUF).id[p]
	}
	return p
}

func (intUF *IntUnionFindV3) isConnected(p, q int) bool {
	return intUF.find(p) == intUF.find(q)
}

func (intUF *IntUnionFindV3) count() int {
	total := 0
	for i, v := range (*intUF).id {
		if i == v {
			total++
		}
	}
	return total
}
