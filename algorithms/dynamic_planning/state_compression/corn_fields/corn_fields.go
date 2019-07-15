package cornfields

const modulo = 100000000

func bitToInt(a []int) int {
	ans := 0
	for i := range a {
		if a[i] == 1 {
			ans += 1 << uint(i)
		}
	}
	return ans
}

func isValid(n int) bool {
	return n&(n<<1) == 0
}

func countWays(farmland [][]int) int {
	rows := make([]int, len(farmland))
	for i := range farmland {
		rows[i] = bitToInt(farmland[i])
	}

	mat := [][]int{}
	for range farmland {
		arr := make([]int, 1<<uint(len(farmland[0])))
		mat = append(mat, arr)
	}

	for i := range rows {
		for j := 0; j <= rows[i]; j++ {
			if isValid(j) && (j|rows[i] == rows[i]) {
				if i == 0 {
					mat[i][j] = 1
				} else {
					for k := 0; k <= rows[i-1]; k++ {
						if mat[i-1][k] > 0 && k&j == 0 {
							mat[i][j] = (mat[i][j] + mat[i-1][k]) % modulo
						}
					}
				}
			}
		}
	}

	ans := 0
	for j := 0; j <= rows[len(rows)-1]; j++ {
		ans = (ans + mat[len(rows)-1][j]) % modulo
	}

	return ans
}
