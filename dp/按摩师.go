package dp

/**
https://leetcode-cn.com/problems/the-masseuse-lcci/comments/
*/
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func massage(nums []int) int {
	size := len(nums)

	if size <= 0 {
		return 0
	}

	// 0代表今天不接单
	// 1代表今天接单
	dp := make([][2]int, size, size)
	dp[0][0] = 0
	dp[0][1] = nums[0]

	for i := 1; i < size; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1])
		dp[i][1] = dp[i-1][0] + nums[i]
	}

	return max(dp[size-1][0], dp[size-1][1])
}
