package hot100

func numIslands(grid [][]byte) int {
	// 染色为-1
	var color func(i, j int)
	color = func(i, j int) {
		if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == '0' {
			return
		}
		// color index
		grid[i][j] = '0'
		color(i + 1, j)
		color(i - 1, j)
		color(i, j + 1)
		color(i, j - 1)
	}

	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				color(i, j)
				res++
			}
		}
	}
	return res
}
