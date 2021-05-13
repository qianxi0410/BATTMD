package hot100

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	row, col := len(matrix), len(matrix[0])
	height := make([]int, col)
	ans := 0

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if matrix[i][j] == '1' {
				height[j] += 1
			} else {
				height[j] = 0
			}
		}
		ans = max(ans, largestRectangleArea(height))
	}
	return ans
}

