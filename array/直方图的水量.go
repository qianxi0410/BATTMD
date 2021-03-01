package main

/**
https://leetcode-cn.com/problems/volume-of-histogram-lcci/
*/

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func trap(height []int) int {
	size := len(height)

	if size <= 2 {
		return 0
	}

	left := make([]int, size)
	right := make([]int, size)

	left[0] = height[0]
	right[size-1] = height[size-1]

	for i := 1; i < size; i++ {
		left[i] = max(left[i-1], height[i])
	}

	for i := size - 2; i >= 0; i-- {
		right[i] = max(right[i+1], height[i])
	}

	res := 0

	for i := 1; i < size-1; i++ {
		tmp := min(left[i-1], right[i+1]) - height[i]
		if tmp > 0 {
			res += tmp
		}
	}

	return res
}
