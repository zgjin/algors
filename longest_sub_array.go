package algors

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = 1

	result := 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}

		result = max(dp[i], result)
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
