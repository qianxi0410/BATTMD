package tree

/**
https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree/submissions/
*/

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if p.Val == root.Val || q.Val == root.Val {
		return root
	}
	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	} else if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	} else {
		return root
	}
}
