package hot100

func swapPairs(head *ListNode) *ListNode {
	cur := head
	pre := &ListNode{}

	for cur != nil {
		next := cur.Next
		if next == nil {
			return head
		}
		tmp := next.Next
		cur.Next = tmp
		next.Next = cur

		if cur == head {
			head = next
		}
		if pre != nil {
			pre.Next = next
		}

		pre = cur
		cur = tmp
	}

	return head
}
