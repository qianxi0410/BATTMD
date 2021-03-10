package tree

/**
https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/
*/
func zigzagLevelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0, 10)

	if root == nil {
		return res
	}
	flag := true
	queue := make([]*TreeNode, 0, 10)
	queue = append(queue, root)

	for len(queue) > 0 {
		size := len(queue)
		tmp := make([]int, size, size)

		for i := 0; i < size; i++ {
			top := queue[0]
			queue = queue[1:]

			if flag {
				tmp[size-i-1] = top.Val
			} else {
				tmp[i] = top.Val
			}

			if top.Left != nil {
				queue = append(queue, top.Left)
			}
			if top.Right != nil {
				queue = append(queue, top.Right)
			}
		}
		flag = !flag
		res = append(res, tmp)
	}
	return res
}
