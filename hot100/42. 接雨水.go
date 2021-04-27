package hot100

func trap(height []int) int {
	l,r := make([]int, len(height)), make([]int, len(height))

	for i := 1; i < len(height); i++ {
		l[i] = max(l[i - 1], height[i - 1])
	}

	for i := len(height) - 2; i >= 0; i-- {
		r[i] = max(r[i + 1], height[i + 1])
	}
	res := 0
	for i := range height {
		if min(r[i], l[i]) - height[i] > 0 {
			res += min(r[i], l[i]) - height[i]
		}
	}

	return res
}
