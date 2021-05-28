package hot100

import "math"

func majorityElement(nums []int) int {
	votes, celebrity := 0, math.MinInt32

	for i := range nums {
		if votes == 0 {
			votes++
			celebrity = nums[i]
		} else if nums[i] == celebrity {
			votes++
		} else {
			votes--
		}
	}

	return celebrity
}
