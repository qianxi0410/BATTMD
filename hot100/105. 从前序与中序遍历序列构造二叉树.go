package hot100

func buildTree(preorder []int, inorder []int) *TreeNode {
	var build func(p_begin, p_end, i_begin, i_end int) *TreeNode

	build = func(p_begin, p_end, i_begin, i_end int) *TreeNode {
		if i_begin > i_end {
			return nil
		}

		rootVal := preorder[p_begin]
		rootIndex := -1

		for i := i_begin; i <= i_end; i++ {
			if inorder[i] == rootVal {
				rootIndex = i
				break
			}
		}

		head := &TreeNode{Val: rootVal}
		head.Left = build(p_begin + 1, p_begin + (rootIndex - i_begin), i_begin, rootIndex - 1)
		head.Right = build(p_begin + (rootIndex - i_begin) + 1, p_end, rootIndex + 1, i_end)

		return head
	}
	return build(0, len(preorder) - 1, 0, len(inorder) - 1)
}
