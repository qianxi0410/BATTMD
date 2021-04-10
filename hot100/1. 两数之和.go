package hot100


func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := range nums {
		m[nums[i]] = i
	}

	for i := range nums {
		j := m[target - nums[i]]
		if i != j && j != 0 {
			return []int{i, j}
		}
	}
	return nil
}