package hot100

func maxProduct(nums []int) int {
	n2 := make([]int, len(nums))
	copy(n2, nums)

	m1 := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i - 1] != 0 {
			nums[i] *= nums[i - 1]
		}
		m1 = max(nums[i], m1)
	}

	m2 := n2[len(n2) - 1]
	for i := len(n2) - 2; i >= 0; i-- {
		if n2[i + 1] != 0 {
			n2[i] *= n2[i + 1]
		}
		m2 = max(n2[i], m2)
	}

	return max(m1, m2)
}
