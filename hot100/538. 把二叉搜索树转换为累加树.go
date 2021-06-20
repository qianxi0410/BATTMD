package hot100

func convertBST(root *TreeNode) *TreeNode {
	total, leftSum := 0, 0

	var dfs func(r *TreeNode)
	dfs = func(r *TreeNode) {
		if r == nil {
			return
		}
		dfs(r.Left)
		total += r.Val
		dfs(r.Right)
	}

	// get total
	dfs(root)

	var helper func(r *TreeNode)

	helper = func(r *TreeNode) {
		if r == nil {
			return
		}
		helper(r.Left)
		r.Val, leftSum = total - leftSum, leftSum + r.Val
		helper(r.Right)
	}

	helper(root)
	return root
}
