package unionfind

import "fmt"

// IntUnionFindV4 ...
type IntUnionFindV4 struct {
	id   []int
	size []int
}

func (intUF *IntUnionFindV4) init(n int) {
	(*intUF).id = make([]int, n)
	(*intUF).size = make([]int, n)
	for i := 0; i < n; i++ {
		(*intUF).id[i] = i
		(*intUF).size[i] = 1
	}
}

func (intUF *IntUnionFindV4) union(p, q int) {
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

func (intUF *IntUnionFindV4) find(p int) int {
	for (*intUF).id[p] != p {
		(*intUF).id[p] = (*intUF).id[(*intUF).id[p]]
		p = (*intUF).id[p]
	}
	return p
}

func (intUF *IntUnionFindV4) isConnected(p, q int) bool {
	return intUF.find(p) == intUF.find(q)
}

func (intUF *IntUnionFindV4) count() int {
	total := 0
	for i, v := range (*intUF).id {
		if i == v {
			total++
		}
	}
	fmt.Println((*intUF).id)
	return total
}
