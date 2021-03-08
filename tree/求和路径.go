package tree

/**
https://leetcode-cn.com/problems/paths-with-sum-lcci/submissions/
*/

func pathSum(root *TreeNode, sum int) int {
	var f func(*TreeNode, []int) int
	f = func(r *TreeNode, a []int) (result int) {
		if r != nil {
			s := make([]int, len(a)+1)
			copy(s, a)
			for i := range s {
				s[i] += r.Val
				if s[i] == sum {
					result++
				}
			}
			result += f(r.Left, s) + f(r.Right, s)
		}
		return
	}
	return f(root, []int{})
}
