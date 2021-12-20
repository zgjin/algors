package algors

// dp(i, j) = min(dp(i - 1, j), dp(i - 1, j - 1), dp(i - 1, j + 1))
func minFallingPathSum(matrix [][]int) int {
	res := 99999999999

	var cache [][]int
	for i := 0; i < len(matrix); i++ {
		mem := make([]int, len(matrix))
		for j := 0; j < len(matrix); j++ {
			mem[j] = 66666
		}
		cache = append(cache, mem)
	}

	for j := 0; j < len(matrix); j++ {
		res = minof2(res, dp(matrix, len(matrix)-1, j, cache))
	}

	return res
}

func dp(matrix [][]int, i, j int, cache [][]int) int {
	if i < 0 || j < 0 || i > len(matrix)-1 || j > len(matrix)-1 {
		return 99999
	}
	if cache[i][j] != 66666 {
		return cache[i][j]
	}

	if i == 0 {
		return matrix[i][j]
	}

	cache[i][j] = matrix[i][j] + minof3(dp(matrix, i-1, j, cache), dp(matrix, i-1, j-1, cache), dp(matrix, i-1, j+1, cache))
	return cache[i][j]
}

func minof2(a, b int) int {
	m := a
	if b < m {
		m = b
	}
	return m
}

func minof3(a, b, c int) int {
	m := a
	if b < m {
		m = b
	}

	if c < m {
		m = c
	}

	return m
}
