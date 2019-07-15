package utils

func fastExp(n, m, p int) int {
	ans := 1
	for m > 0 {
		if m&1 == 1 {
			ans = ans * n % p
			m--
		}
		m >>= 1
		n = n * n % p
	}
	return ans
}
