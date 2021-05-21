package hot100

func flatten(root *TreeNode)  {
	var last *TreeNode
	var helper func(r *TreeNode)
	helper = func(r *TreeNode) {
		if r == nil {
			return
		}
		helper(r.Right)
		helper(r.Left)
		r.Right = last
		r.Left = nil
		last = r
	}
	helper(root)
}
