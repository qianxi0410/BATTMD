package hot100

func exist(board [][]byte, word string) bool {
	row, col := len(board), len(board[0])

	dp := make([][]bool, row)
	for i := range dp {
		dp[i] = make([]bool, col)
	}

	var bfs func(i, j int, w string) bool
	bfs = func(i, j int, w string) bool {
		if i < 0 || i >= row || j < 0 || j >= col || board[i][j] != w[0] || dp[i][j] {
			return false
		}
		if len(w) == 1 {
			return true
		}

		dp[i][j] = true

		res := bfs(i + 1, j, w[1:]) || bfs(i, j + 1, w[1:]) || bfs(i - 1, j, w[1:]) || bfs(i, j - 1, w[1:])

		dp[i][j] = false
		return res
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if bfs(i, j, word) {
				return true
			}
		}
	}

	return false
}
