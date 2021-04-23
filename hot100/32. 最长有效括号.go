package hot100

func longestValidParentheses(s string) int {
	stack := make([]int, 0, len(s))
	mark := make([]int, len(s), len(s))
	res := 0

	for i, v := range s {
		switch v {
		case '(':
			stack = append(stack, i)
		case ')':
			if len(stack) > 0 {
				stack = stack[:len(stack) - 1]
			} else {
				mark[i] = 1
			}
		}
	}

	for len(stack) > 0 {
		mark[stack[len(stack) - 1]] = 1
		stack = stack[:len(stack) - 1]
	}

	tmp := 0
	for i := 0; i < len(mark); i++ {
		if mark[i] == 0 {
			tmp++
		} else {
			res = max(res, tmp)
			tmp = 0
		}
	}

	return max(res, tmp)
}