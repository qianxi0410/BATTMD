package hot100

import (
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := make([][]int, 0)
	tmp := make([]int, 0)

	var f func([]int, int, int)
	f = func(tmp []int, sum int, begin int) {
		if sum > target {
			return
		}

		if sum == target {
			t := make([]int, len(tmp))
			copy(t, tmp)
			res = append(res, t)
			return
		}
		for i := begin; i < len(candidates); i++ {
			tmp = append(tmp, candidates[i])
			f(tmp, sum + candidates[i], i)
			tmp = tmp[:len(tmp) - 1]
		}
	}
	f(tmp, 0, 0)
	return res
}