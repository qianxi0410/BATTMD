package hot100

func largestRectangleArea(heights []int) int {
	area := 0
	stk := make([]int, 0, len(heights))

	for i := 0; i < len(heights); i++ {
		for len(stk) > 0 && heights[i] <= heights[stk[len(stk) - 1]] {
			top := stk[len(stk) - 1]
			stk = stk[:len(stk) - 1]
			width := 0

			if len(stk) == 0 {
				width = i
			} else {
				width = i - stk[len(stk) - 1] - 1
			}
			area = max(area, heights[top] * width)
		}
		stk = append(stk, i)

	}

	for len(stk) > 0 {
		top := stk[len(stk) - 1]
		stk = stk[:len(stk) - 1]
		width := 0

		if len(stk) == 0 {
			width = len(heights)
		} else {
			width = len(heights) - stk[len(stk) - 1] - 1
		}
		area = max(area, heights[top] * width)
	}

	return area
}
