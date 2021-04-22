package hot100

import "sort"

func nextPermutation(nums []int)  {
	i := len(nums) - 1
	for i > 0 && nums[i - 1] >= nums[i] {
		i--
	}
	if i == 0 {
		sort.Ints(nums)
	} else {
		j := len(nums) - 1
		for j > i && nums[j] <= nums[i - 1] {
			j--
		}
		nums[i - 1], nums[j] = nums[j], nums[i - 1]
		sort.Ints(nums[i:])
	}
}
