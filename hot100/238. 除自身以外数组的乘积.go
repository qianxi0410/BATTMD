package hot100

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	for i := range res {
		res[i] = 1
	}

	l, r := 1, 1
	for i := range nums {
		res[i] *= l
		l *= nums[i]

		res[len(nums) - 1 - i] *= r
		r *= nums[len(nums) - i - 1]
	}

	return res
}
