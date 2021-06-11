package hot100


func diameterOfBinaryTree(root *TreeNode) int {
	var res = 0

	var f func(r *TreeNode) int
	f = func(r *TreeNode) int {
		if r == nil {
			return 0
		}
		return max(f(r.Left), f(r.Right))
	}

	var dfs func(r *TreeNode)
	dfs = func(r *TreeNode) {
		if r == nil {
			return
		}
		dfs(r.Left)
		res = max(res, f(r.Right) + f(r.Left))
		dfs(r.Right)
	}

	dfs(root)

	return res
}