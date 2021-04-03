package math

type ListNode struct {
     Val int
     Next *ListNode
}

/**
https://leetcode-cn.com/problems/sum-lists-lcci/
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 进位
	cur := 0
	head := &ListNode{
		Val:  -1,
		Next: nil,
	}
	tmpHead := head

	for l1 != nil || l2 != nil {
		tmp := cur
		if l1 != nil {
			tmp += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			tmp += l2.Val
			l2 = l2.Next
		}

		if tmp >= 10 {
			cur = tmp / 10
			tmp -= cur * 10
		} else {
			cur = 0
		}

		head.Next = &ListNode{
			Val:  tmp,
			Next: nil,
		}
		head = head.Next
	}

	if cur != 0 {
		head.Next = &ListNode{
			Val:  1,
			Next: nil,
		}
	}
	return tmpHead.Next
}
