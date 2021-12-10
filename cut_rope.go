package algors

import "math"

// f(n) = f(i)* f(n - i)
// f(2) = 1
// f(3) = 2
// f(4) = 4
// Math.max(j*(i-j),j*dp[i-j]是由于减去第一段长度为j的绳子后，可以继续剪也可以不剪
// Math.max(dp[i],Math.max(j*(i-j),j*dp[i-j]))是当j不同时，求出最大的dp[i]
func cuttingRope(n int) int {
	if n == 2 || n == 1 {
		return 1
	}
	if n == 3 {
		return 2
	}
	if n == 4 {
		return 4
	}

	cache := make([]int, n+1)
	cache[2] = 1
	cache[2] = 1
	cache[3] = 2
	cache[4] = 4

	for i := 5; i <= n; i++ {
		for j := 1; j < i; j++ {
			cache[i] = int(math.Max(float64(cache[i]), math.Max(float64(j*(i-j)), float64(j*cache[i-j]))))
		}
	}

	return cache[n]
}
