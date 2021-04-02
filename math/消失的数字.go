package math

/**
https://leetcode-cn.com/problems/missing-number-lcci/comments/
 */

func missingNumber(nums []int) int {
	n := len(nums)
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return (n * (n - 1)) / 2 - sum
}
