package segmenttree

type segTree struct {
	t []int64
	d []int64
	n int64
	h int64
}

func newSegTree(nums []int64) segTree {
	n := len(nums)
	t := make([]int64, n*2)
	d := make([]int64, n)
	for i := range nums {
		t[n+i] = nums[i]
	}
	for i := n - 1; i > 0; i-- {
		t[i] = t[i<<1] + t[i<<1|1]
	}
	h := 0
	for n > 0 {
		n >>= 1
		h++
	}
	return segTree{
		t: t,
		d: d,
		n: int64(len(nums)),
		h: int64(h),
	}
}

func (s segTree) calc(p, k int64) {
	if s.d[p] == 0 {
		s.t[p] = s.t[p<<1] + s.t[p<<1|1]
	} else {
		s.t[p] += s.d[p] * k
	}
}

func (s segTree) apply(p, val, k int64) {
	s.t[p] += val * k
	if p < s.n {
		s.d[p] += val
	}
}

func (s segTree) push(l, r int64) {
	j := s.h
	k := int64(1) << uint(j-1)
	for l, r = l+s.n, r+s.n-1; j > 0; j, k = j-1, k>>1 {
		for i := l >> uint(j); i <= r>>uint(j); i++ {
			if s.d[i] != 0 {
				s.apply(i<<1, s.d[i], k)
				s.apply(i<<1|1, s.d[i], k)
				s.d[i] = 0
			}
		}
	}
}

func (s segTree) modify(l, r, val int64) {
	if val == 0 {
		return
	}
	s.push(l, l+1)
	s.push(r-1, r)
	cl := false
	cr := false
	k := int64(1)
	for l, r = l+s.n, r+s.n; l < r; l, r, k = l>>1, r>>1, k<<1 {
		if cl {
			s.calc(l-1, k)
		}
		if cr {
			s.calc(r, k)
		}
		if l&1 == 1 {
			s.apply(l, val, k)
			l++
			cl = true
		}
		if r&1 == 1 {
			r--
			s.apply(r, val, k)
			cr = true
		}
	}

	for l--; r > 0; l, r, k = l>>1, r>>1, k<<1 {
		if cl {
			s.calc(l, k)
		}
		if cr && (!cl || l != r) {
			s.calc(r, k)
		}
	}
}

func (s segTree) query(l, r int64) int64 {
	s.push(l, l+1)
	s.push(r-1, r)
	ans := int64(0)
	for l, r = l+s.n, r+s.n; l < r; l, r = l>>1, r>>1 {
		if l&1 == 1 {
			ans += s.t[l]
			l++
		}
		if r&1 == 1 {
			r--
			ans += s.t[r]
		}
	}
	return ans
}
