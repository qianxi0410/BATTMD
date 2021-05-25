package hot100

func longestConsecutive(nums []int) int {
	m := make(map[int]int, len(nums))

	for i := 0; i < len(nums); i++ {
		m[nums[i]] = 1
	}

	res := 0
	for k := range m {
		i, sum := k, 0
		for m[i] != 0 {
			sum += m[i]
			m[i] = 0
			i--
		}
		m[k] = sum
		res = max(res, sum)
	}

	return res
}
