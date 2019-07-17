package unionfind

// IntUnionFindV1 ...
type IntUnionFindV1 struct {
	id []int
}

func (intUF *IntUnionFindV1) init(n int) {
	(*intUF).id = make([]int, n)
	for i := 0; i < n; i++ {
		(*intUF).id[i] = i
	}
}

func (intUF *IntUnionFindV1) union(p, q int) {
	source := (*intUF).id[p]
	target := (*intUF).id[q]
	for i := range (*intUF).id {
		if (*intUF).id[i] == source {
			(*intUF).id[i] = target
		}
	}
}

func (intUF *IntUnionFindV1) find(p int) int { return (*intUF).id[p] }

func (intUF *IntUnionFindV1) isConnected(p, q int) bool {
	return (*intUF).id[p] == (*intUF).id[q]
}

func (intUF *IntUnionFindV1) count() int {
	occurred := map[int]bool{}
	total := 0
	for _, v := range (*intUF).id {
		if !occurred[v] {
			total++
			occurred[v] = true
		}
	}
	return total
}
