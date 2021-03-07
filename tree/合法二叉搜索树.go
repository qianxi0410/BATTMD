package tree

/**
https://leetcode-cn.com/problems/legal-binary-search-tree-lcci/comments/
*/

func isValidBST(root *TreeNode) bool {
	return isValidate(root, -0x3f3f3f3f, 0x3f3f3f3f)
}

func isValidate(root *TreeNode, min, max int) bool {
	return root == nil || (root.Val > min && root.Val < max) && isValidate(root.Left, min, root.Val) && isValidate(root.Right, root.Val, max)
}
