package algors

// dp[i][j] = 0 (obstacleGrid[i][j] = 1)
// dp[i][j] = dp[i-1][j] + dp[i-1][j]
// dp[i][0] = 1
// dp[0][j] = 1
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	dp := make([][]int, m)
	var ib bool
	var jb bool
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			if i == 0 {
				if obstacleGrid[i][j] != 1 && !ib {
					dp[i][j] = 1
				} else {
					ib = true
				}
			}

			if j == 0 {
				if obstacleGrid[i][j] != 1 && !jb {
					dp[i][j] = 1
				} else {
					jb = true
				}
			}
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}

	return dp[m-1][n-1]
}
