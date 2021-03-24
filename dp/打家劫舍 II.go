package dp

import "math"
/**
https://leetcode-cn.com/problems/house-robber-ii/comments/
 */
func maxInts(arr []int, ans int) int {
	res := -1
	for i := 0; i < len(arr); i++ {
		res = max(res, arr[i])
	}
	return max(res, ans)
}

func rob(nums []int) int {
	size := len(nums)
	if size == 1 {
		return nums[0]
	}

	// 4种状态
	// 0 -> 偷i偷0
	// 1 -> 偷i不偷0
	// 2 -> 不偷i偷0
	// 3 -> 不偷i不偷0
	dp := make([][4]int, size)
	dp[0][0] = nums[0]
	dp[0][1] = math.MinInt32
	dp[0][2] = math.MinInt32
	dp[0][3] = 0

	res := math.MinInt32
	for i := 1; i < size; i++ {
		dp[i][0] = dp[i - 1][2] + nums[i]
		dp[i][1] = dp[i - 1][3] + nums[i]
		dp[i][2] = max(dp[i - 1][0], dp[i - 1][2])
		dp[i][3] = max(dp[i - 1][1], dp[i - 1][3])
		if i == size - 1 {
			res = maxInts(dp[i][1:], res)
		}
	}
	return res
}