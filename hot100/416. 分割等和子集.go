package hot100

func canPartition(nums []int) bool {
	sum := 0

	for i := range nums {
		sum += nums[i]
	}
	if sum%2 != 0 {
		return false
	}
	target := sum / 2
	dp := make([]int, target+1)
	dp[0] = 1
	for _, n := range nums {
		for i := target; i >= n; i-- {
			dp[i] += dp[i-n]
		}
	}
	return dp[target] != 0
}
