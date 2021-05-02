package hot100

import "math"

func maxSubArray(nums []int) int {
	res, sum := math.MinInt32, 0

	for i := range nums {
		sum += nums[i]
		if sum >= 0 {
			nums[i] = sum
		} else {
			sum = 0
		}
	}

	for i := range nums {
		res = max(res, nums[i])
	}

	return res
}
