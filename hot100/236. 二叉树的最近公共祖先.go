package hot100

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == p || root == q || root == nil {
		return root
	}
	l, r := lowestCommonAncestor(root.Left, p, q), lowestCommonAncestor(root.Right, p, q)
	if l != nil && r != nil {
		return root
	} else if l == nil || r == nil {
		if l != nil {
			return l
		} else {
			return r
		}
	}
	return nil
}