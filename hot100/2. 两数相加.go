package hot100

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{
		Val:  -1,
		Next: nil,
	}
	p := head

	carry := 0

	for l1 != nil || l2 != nil {
		tmp := carry
		if l1 != nil {
			tmp += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			tmp += l2.Val
			l2 = l2.Next
		}

		carry = tmp / 10
		tmp = tmp % 10
		p.Next = &ListNode{Val: tmp}
		p = p.Next
	}

	if carry != 0 {
		p.Next = &ListNode{Val: carry}
	}
	return head.Next
}
