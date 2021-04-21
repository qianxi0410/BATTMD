package hot100

import "math"

func mergeKLists(lists []*ListNode) *ListNode {
	k := len(lists)
	head := &ListNode{}
	p := head

	for {
		flag := false
		minVal := math.MaxInt32
		var minIndex int
		for i := 0; i < k; i++ {
			if lists[i] != nil && lists[i].Val < minVal {
				minVal = lists[i].Val
				minIndex = i
				flag = true
			}
		}
		if flag {
			p.Next = lists[minIndex]
			p = p.Next
			lists[minIndex] = lists[minIndex].Next
		} else {
			break
		}
	}
	p.Next = nil
	return head.Next
}
