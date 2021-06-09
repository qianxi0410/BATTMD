package hot100

func findDisappearedNumbers(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		if nums[abs(nums[i]) - 1] > 0 {
			nums[abs(nums[i]) - 1] = -nums[abs(nums[i]) - 1]
		}
	}

	res := make([]int, 0, len(nums))
	for i := 1; i <= len(nums); i++ {
		if nums[i - 1] > 0 {
			res = append(res, i)
		}
	}
	return res
}