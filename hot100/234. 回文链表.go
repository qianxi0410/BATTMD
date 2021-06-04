package hot100

func isPalindrome(head *ListNode) bool {
	stk := make([]int, 0)

	p := head
	for p != nil {
		stk = append(stk, p.Val)
		p = p.Next
	}
	p = head
	for len(stk) > 0 {
		top := stk[len(stk) - 1]
		stk = stk[:len(stk) - 1]
		if top != p.Val {
			return false
		}
		p = p.Next
	}
	return true
}