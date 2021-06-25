package hot100

func lengthOfLIS(nums []int) int {
	size, res := len(nums), make([]int, len(nums))

	for i := range res {
		res[i] = 1
	}
	ans := 1
	for i := 1; i < size; i++ {
		for j := i - 1; j >= 0; j-- {
			if nums[i] > nums[j] {
				res[i] = max(res[j] + 1, res[i])
			}
		}
		ans = max(ans, res[i])
	}
	return ans
}