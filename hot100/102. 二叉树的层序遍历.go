package hot100

func levelOrder(root *TreeNode) [][]int {
	q := make([]*TreeNode, 0)
	res := make([][]int, 0)

	if root == nil {
		return res
	}

	q = append(q, root)

	for len(q) != 0 {
		size := len(q)
		tmp := make([]int, 0, size)
		for i := 0; i < size; i++ {
			front := q[0]
			tmp = append(tmp, front.Val)
			q = q[1:]
			if front.Left != nil {
				q = append(q, front.Left)
			}
			if front.Right != nil {
				q = append(q, front.Right)
			}
		}
		res = append(res, tmp)
	}
	return res
}
