package tree

/**
https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-di-kda-jie-dian-lcof/
*/
func inorder(root *TreeNode, res []int) {
	if root == nil {
		return
	}
	inorder(root.Left, res)
	res = append(res, root.Val)
	inorder(root.Right, res)
}

func kthLargest(root *TreeNode, k int) int {
	res := make([]int, 0)
	inorder(root, res)
	return res[len(res)-1-k]
}
