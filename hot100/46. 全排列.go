package hot100

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	size := len(nums)

	var f func([]int, int)
	f = func(tmp []int, curNum int) {
		if len(tmp) == size {
			t := make([]int, len(tmp))
			copy(t, tmp)
			res = append(res, t)
			return
		}

		for i := 0; i < size; i++ {
			if find(tmp, nums[i]) {
				continue
			}

			tmp = append(tmp, nums[i])
			f(tmp, nums[i])
			tmp = tmp[:len(tmp) - 1]
		}
	}
	tmp := make([]int, 0, len(nums))
	f(tmp, 0)
	return res
}

func find(arr []int, num int) bool {
	for i := range arr {
		if arr[i] == num {
			return true
		}
	}
	return false
}