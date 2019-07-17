package unionfind

// UnionFind ...
type UnionFind interface {
	init(n int)
	union(p, q int)
	find(p int) int
	isConnected(p, q int) bool
	count() int
}
