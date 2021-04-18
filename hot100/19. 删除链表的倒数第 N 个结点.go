package hot100

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	newHead := &ListNode{Next: head}

	slow, fast := head, head
	mark := newHead

	for i := 0; i < n && fast != nil; i++ {
		fast = fast.Next
	}

	for fast != nil {
		mark = mark.Next
		slow = slow.Next
		fast = fast.Next
	}
	mark.Next = slow.Next

	return newHead.Next
}
