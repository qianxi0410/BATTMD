package hot100

func isValid(s string) bool {
	stk := make([]rune, 0)

	for _, v := range s {
		switch v {
		case '(', '{', '[':
			stk = append(stk, v)
		case ')':
			if len(stk) == 0 || stk[len(stk) - 1] != '(' {
				return false
			}
			stk = stk[:len(stk) - 1]
		case '}':
			if len(stk) == 0 || stk[len(stk) - 1] != '{' {
				return false
			}
			stk = stk[:len(stk) - 1]
		case ']':
			if len(stk) == 0 || stk[len(stk) - 1] != '[' {
				return false
			}
			stk = stk[:len(stk) - 1]
		}
	}
	if len(stk) == 0 {
		return true
	}
	return false
}