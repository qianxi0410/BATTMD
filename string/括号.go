package string

/**
https://leetcode-cn.com/problems/bracket-lcci/comments/
 */
func generateParenthesis(n int) []string {
	res := make([]string, 0)

	var f func(left, right int, track string)
	f = func(left, right int,  track string) {
		if left < right || left > n || right > n {
			return
		} else if left == n && right == n {
			res = append(res, track)
			return
		} else {
			f(left + 1, right, track + "(")
			f(left, right + 1, track + ")")
		}
	}
	f(0, 0, "")
	return res
}