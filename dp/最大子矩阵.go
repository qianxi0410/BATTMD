package dp

import "math"

/**
https://leetcode-cn.com/problems/max-submatrix-lcci/submissions/
*/

func getMaxMatrix(matrix [][]int) []int {
	n, m := len(matrix), len(matrix[0])
	max, ans := math.MinInt32, make([]int, 4)
	for i := 0; i < n; i++ {
		colSum := make([]int, m)
		for j := i; j < n; j++ {
			for k := 0; k < m; k++ {
				colSum[k] += matrix[j][k] // 计算截止到当前行的列的和
			}

			// 简化为计算最大子数组
			sum, start := 0, 0
			for k := 0; k < m; k++ {
				if sum <= 0 {
					start = k
					sum = 0
				}
				sum += colSum[k]
				if sum > max {
					max = sum
					ans = []int{i, start, j, k}
				}
			}
		}
	}
	return ans
}
