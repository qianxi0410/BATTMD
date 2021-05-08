package hot100

func minDistance(word1 string, word2 string) int {
	row, col := len(word1), len(word2)

	if row == 0 {
		return col
	}

	if col == 0 {
		return row
	}

	dp := make([][]int, row + 1)
	for i := range dp {
		dp[i] = make([]int, col + 1)
	}

	for i := 0; i <= row; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= col; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= row; i++ {
		for j := 1; j <= col; j++ {
			if word1[i - 1] == word2[j - 1] {
				dp[i][j] = dp[i - 1][j - 1]
			} else {
				dp[i][j] = min(min(dp[i - 1][j], dp[i][j - 1]), dp[i - 1][j - 1]) + 1
			}
		}
	}

	return dp[row][col]
}
