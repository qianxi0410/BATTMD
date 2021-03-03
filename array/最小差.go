package main

import "sort"

/**
https://leetcode-cn.com/problems/smallest-difference-lcci/
*/

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func smallestDifference(a []int, b []int) int {
	sizea := len(a)
	sizeb := len(b)

	sort.Ints(a)
	sort.Ints(b)

	res := int(^uint(0) >> 1)

	i, j := 0, 0

	for i < sizea && j < sizeb {
		diff := a[i] - b[j]
		res = min(abs(diff), res)
		if diff < 0 {
			i++
		} else {
			j++
		}
	}
	return res
}
