package hot100

func inorderTraversal(root *TreeNode) []int {
	stk := make([]*TreeNode, 0)
	res := make([]int, 0)
	cur := root

	for cur != nil || len(stk) > 0 {
		if cur != nil {
			stk = append(stk, cur)
			cur = cur.Left
		} else {
			cur = stk[len(stk) - 1]
			stk = stk[:len(stk) - 1]
			res = append(res, cur.Val)
			cur = cur.Right
		}
	}
	return res
}
