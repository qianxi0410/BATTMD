package math

/**
https://leetcode-cn.com/problems/factorial-zeros-lcci/comments/
 */

func trailingZeroes(n int) int {
	cnt := 0
	for n >= 5 {
		n /= 5
		cnt += n
	}
	return cnt
}
