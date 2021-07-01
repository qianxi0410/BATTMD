package hot100

func maximalSquare(matrix [][]byte) int {
	res := 0

	row, col := len(matrix), len(matrix[0])
	if row < 1 {
		return 0
	}
	dp := make([][]int, row+1)
	for i := range dp {
		dp[i] = make([]int, col+1)
	}

	for i := 1; i <= row; i++ {
		for j := 1; j <= col; j++ {
			if matrix[i-1][j-1] == '1' {
				dp[i][j] = 1 + min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1]))
				res = max(res, dp[i][j])
			}
		}
	}
	return res * res
}
