package hot100

func _cut(head *ListNode, n int) *ListNode {
	p := head
	for i := n - 1; i > 0 && p != nil; i-- {
		p = p.Next
	}

	if p == nil {
		return nil
	}
	next := p.Next
	p.Next = nil
	return next
}

func _merge(l1, l2 *ListNode) *ListNode {
	tmpHead := &ListNode{}
	p := tmpHead

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			p.Next = l1
			p = p.Next
			l1 = l1.Next
		} else {
			p.Next = l2
			p = p.Next
			l2 = l2.Next
		}
	}

	if l1 != nil {
		p.Next = l1
	} else {
		p.Next = l2
	}

	return tmpHead.Next
}

func sortList(head *ListNode) *ListNode {
	tmpHead := &ListNode{Next: head}
	p := head
	length := 0

	for p != nil {
		length++
		p = p.Next
	}

	for size := 1; size < length; size <<= 1 {
		cur, tail := tmpHead.Next, tmpHead
		for cur != nil {
			left := cur
			right := _cut(left, size)
			cur = _cut(right, size)

			tail.Next = _merge(left, right)
			for tail.Next != nil {
				tail = tail.Next
			}
		}
	}
	return tmpHead.Next
}
