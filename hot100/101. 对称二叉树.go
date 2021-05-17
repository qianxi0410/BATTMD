package hot100

func isSymmetric(root *TreeNode) bool {
	var foo func(l, r *TreeNode) bool
	foo = func(l, r *TreeNode) bool {
		if l == nil && r == nil {
			return true
		} else if l == nil || r == nil {
			return false
		} else {
			return l.Val == r.Val && foo(l.Left, r.Right) && foo(l.Right, r.Left)
		}
	}
	return foo(root, root)
}
