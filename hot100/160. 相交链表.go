package hot100


func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lenA, lenB := 0, 0
	tmpA, tmpB := headA, headB
	for tmpA != nil {
		lenA++
		tmpA = tmpA.Next
	}
	for tmpB != nil {
		lenB++
		tmpB = tmpB.Next
	}

	var long *ListNode
	var short *ListNode

	if lenA > lenB {
		long = headA
		short = headB
	} else {
		long = headB
		short = headA
	}

	for i := 0; i < abs(lenA - lenB); i++ {
		long = long.Next
	}
	for long != nil && short != nil {
		if long == short {
			return long
		}
		long = long.Next
		short = short.Next
	}
	return nil
}
