package main

/**
https://leetcode-cn.com/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof/
*/

func findNumberIn2DArray(matrix [][]int, target int) bool {
	row := len(matrix)
	if row == 0 {
		return false
	}
	col := len(matrix[0])
	if col == 0 {
		return false
	}

	i, j := 0, col-1

	for {
		if target > matrix[i][j] {
			i++
			if i >= row {
				break
			}
		} else if target < matrix[i][j] {
			j--
			if j < 0 {
				break
			}
		} else {
			return true
		}
	}

	return false
}
