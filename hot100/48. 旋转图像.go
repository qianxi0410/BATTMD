package hot100

func rotate(matrix [][]int)  {
	reverse := func(arr []int) {
		for i := 0; i < len(arr) / 2; i++ {
			arr[i], arr[len(arr) - 1 - i] = arr[len(arr) - 1 - i], arr[i]
		}
	}

	fold := func(i, j int) {
		row := len(matrix)
		matrix[i][j], matrix[row - 1 - j][row - 1 - i] = matrix[row - 1 - j][row - 1 - i], matrix[i][j]
	}

	for i := range matrix {
		reverse(matrix[i])
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix) - i; j++ {
			fold(i, j)
		}
	}
}
