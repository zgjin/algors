package algors

// 获取字符串 s 中最长有效的括号序列
// dp[i] 以 i 为结尾的最长有效序列
// )()()) ())((()))
// dp[i] = dp[i-2] +2 (s[i-1] == '(')
// dp[i] = dp[i−1]+dp[i−dp[i−1]−2]+2 (s[i-1] = ')' && s[i - dp[i-1] -1] == '(')
func longestValidParentheses(s string) int {
	dp := make([]int, len(s))

	var max int

	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				if i < 2 {
					dp[i] = 2
				} else {
					dp[i] = dp[i-2] + 2
				}
			} else if s[i-1] == ')' && i-dp[i-1] > 0 && s[i-dp[i-1]-1] == '(' {
				if i-dp[i-1] >= 2 {
					dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2
				} else {
					dp[i] = dp[i-1] + 2
				}
			}
		}

		if dp[i] > max {
			max = dp[i]
		}
	}

	return max
}
