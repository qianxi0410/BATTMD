package hot100

func maxArea(height []int) int {
	l, r := 0, len(height) - 1
	res := -1

	for l < r {
		short := min(height[l], height[r])
		res = max(res, (r - l) * short)
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return res
}
