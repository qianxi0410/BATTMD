package hot100

func findTargetSumWays(nums []int, target int) int {
	res, size := 0, len(nums)

	var dfs func(i, sum int)
	dfs = func(i, sum int) {
		if i == size - 1 && sum == target {
			res++
		} else {
			if i >= size - 1 {
				return
			}
			dfs(i + 1, sum + nums[i + 1])
			dfs(i + 1, sum - nums[i + 1])
		}
	}

	dfs(0, nums[0])
	dfs(0, -nums[0])
	return res
}
