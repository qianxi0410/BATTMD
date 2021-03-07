package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
https://leetcode-cn.com/problems/list-of-depth-lcci/
*/

func listOfDepth(tree *TreeNode) []*ListNode {
	res := make([]*ListNode, 0)
	queue := make([]*TreeNode, 0)

	if tree == nil {
		return res
	}

	queue = append(queue, tree)
	for len(queue) > 0 {
		size := len(queue)
		virtualHead := &ListNode{}
		pre := virtualHead
		for i := 0; i < size; i++ {
			top := queue[0]
			queue = queue[1:]
			node := &ListNode{
				Val:  top.Val,
				Next: nil,
			}
			pre.Next = node
			pre = pre.Next

			if top.Left != nil {
				queue = append(queue, top.Left)
			}
			if top.Right != nil {
				queue = append(queue, top.Right)
			}
		}
		res = append(res, virtualHead.Next)
	}

	return res
}
