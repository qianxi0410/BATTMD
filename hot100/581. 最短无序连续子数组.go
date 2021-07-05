package hot100

import "sort"

func findUnsortedSubarray(nums []int) int {
	tmp := make([]int, len(nums))
	copy(tmp, nums)
	sort.Ints(tmp)
	begin, end := -1, -1
	for i := range nums {
		if nums[i] != tmp[i] && begin == -1 {
			begin = i
		} else if nums[i] != tmp[i] {
			end = i
		}
	}
	if end == begin {
		return 0
	}

	return end - begin + 1
}
