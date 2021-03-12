package tree

/**
https://leetcode-cn.com/problems/dui-cheng-de-er-cha-shu-lcof/
*/
func helper(r1, r2 *TreeNode) bool {
	if r1 == nil && r2 == nil {
		return true
	} else if r1 == nil || r2 == nil {
		return false
	}
	return r1.Val == r2.Val && helper(r1.Left, r2.Right) && helper(r1.Right, r2.Left)
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return helper(root.Left, root.Right)
}
