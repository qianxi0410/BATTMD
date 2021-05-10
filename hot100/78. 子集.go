package hot100

func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	var f func(begin int, tmp []int)

	f = func(begin int, tmp []int) {
		if len(tmp) <= len(nums) {
			t := make([]int, len(tmp))
			copy(t, tmp)
			res = append(res, t)
		}

		for i := begin; i < len(nums); i++ {
			tmp = append(tmp, nums[i])
			f(i + 1, tmp)
			tmp = tmp[: len(tmp) - 1]
		}
	}
	tmp := make([]int, 0, len(nums))
	f(0, tmp)
	return res
}