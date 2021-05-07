package hot100

func climbStairs(n int) int {
	dp := make([]int, n)
	if n <= 2 {
		return n
	}
	dp[0], dp[1] = 1, 2

	for i := 3; i < n; i++ {
		dp[i] = dp[i - 1] + dp[i - 2]
	}

	return dp[n - 1]
}
