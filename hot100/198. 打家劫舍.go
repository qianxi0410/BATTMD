package hot100

func rob(nums []int) int {
	dp := make([][2]int, len(nums))
	dp[0][0], dp[0][1] = 0, nums[0]

	for i := 1; i < len(nums); i++ {
		dp[i][0] = max(dp[i - 1][0], dp[i - 1][1])
		dp[i][1] = dp[i - 1][0] + nums[i]
	}

	return max(dp[len(nums) - 1][0], dp[len(nums) - 1][1])
}
