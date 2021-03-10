package tree

/**
https://leetcode-cn.com/problems/shu-de-zi-jie-gou-lcof/
*/

func dfs(A, B *TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil {
		return false
	}
	return A.Val == B.Val && dfs(A.Left, B.Left) && dfs(A.Right, B.Right)
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}

	return dfs(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}
