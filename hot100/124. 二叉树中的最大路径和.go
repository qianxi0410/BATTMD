package hot100

import "math"

func maxPathSum(root *TreeNode) int {
	res := math.MinInt32

	var getMax func(root *TreeNode) int
	getMax = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		l, r := max(0, getMax(root.Left)), max(0, getMax(root.Right))
		res = max(res, root.Val+l+r)
		return max(l, r) + root.Val
	}
	getMax(root)
	return res
}
