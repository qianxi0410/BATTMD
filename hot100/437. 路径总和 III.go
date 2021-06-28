package hot100

func pathSum(root *TreeNode, targetSum int) int {
	cnt := 0

	var sum func(r *TreeNode, tmpSum int)
	sum = func(r *TreeNode, tmpSum int) {
		if r == nil {
			return
		}
		if tmpSum + r.Val == targetSum {
			cnt++
		}
		sum(r.Left, tmpSum + r.Val)
		sum(r.Right, tmpSum + r.Val)
	}

	var inOrder func(r *TreeNode)
	inOrder = func(r *TreeNode) {
		if r == nil {
			return
		}
		inOrder(r.Left)
		sum(r, 0)
		inOrder(r.Right)
	}
	inOrder(root)
	return cnt
}