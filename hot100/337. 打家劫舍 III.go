package hot100

func _rob(root *TreeNode) int {
	var helper func(r *TreeNode) []int
	helper = func(root *TreeNode) []int {
		if root == nil {
			return make([]int, 2)
		}
		l, r := helper(root.Left), helper(root.Right)
		// res[0]表示不抢该节点可获得最大值,res[1]表示抢劫该节点可获得最大值
		res := make([]int, 2)
		res[0], res[1] = max(l[0], l[1])+max(r[0], r[1]), root.Val+l[0]+r[0]
		return res
	}
	helper(root)

	res := helper(root)
	return max(res[0], res[1])
}
