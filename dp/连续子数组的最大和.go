package dp

/**
https://leetcode-cn.com/problems/lian-xu-zi-shu-zu-de-zui-da-he-lcof/comments/
*/
func maxSubArray(nums []int) int {
	size := len(nums)
	if size <= 0 {
		return 0
	}

	if size == 1 {
		return nums[0]
	}

	res := 0
	dp := make([]int, size)

	dp[0] = nums[0]
	res = nums[0]

	for i := 1; i < size; i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		res = max(dp[i], res)
	}

	return res
}
