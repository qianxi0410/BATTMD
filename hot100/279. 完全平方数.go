package hot100

import "math"

func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		minVal := math.MaxInt32
		for j := 1; j*j <= i; j++ {
			minVal = min(minVal, dp[i-j*j])
			dp[i] = minVal + 1
		}
	}
	return dp[n]
}
