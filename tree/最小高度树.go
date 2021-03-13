package tree

import "sort"

func build(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}

	mid := left + (right-left)>>1
	head := &TreeNode{
		Val:   nums[mid],
		Left:  nil,
		Right: nil,
	}
	head.Left = build(nums, left, mid-1)
	head.Right = build(nums, mid+1, right)
	return head
}

func sortedArrayToBST(nums []int) *TreeNode {
	sort.Ints(nums)
	return build(nums, 0, len(nums))
}
