package hot100

import "math"

func isValidate(root *TreeNode, min, max int) bool {
	return root == nil || (root.Val > min && root.Val < max) && isValidate(root.Left, min, root.Val) && isValidate(root.Right, root.Val, max)
}

func isValidBST(root *TreeNode) bool {
	return isValidate(root, math.MinInt32, math.MaxInt32)
}