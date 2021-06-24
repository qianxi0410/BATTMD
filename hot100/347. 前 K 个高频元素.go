package hot100

import "sort"

type pair struct {
	val int
	count int
}

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	pairs := make([]pair, 0, len(nums))

	for i := range nums {
		m[nums[i]]++
	}

	for k, v := range m {
		pairs = append(pairs, pair{
			val:   k,
			count: v,
		})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].count > pairs[j].count
	})

	res := make([]int, 0, k)
	for i := 0; i < k; i++ {
		res = append(res, pairs[i].val)
	}
	return res
}
