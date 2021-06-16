package hot100

func countSubstrings(s string) int {
	res := 0

	var count func(begin, end int)
	count = func(begin, end int) {
		for begin >= 0 && end < len(s) && s[begin] == s[end] {
			res++
			begin--
			end++
		}
	}

	for i := range s {
		count(i, i)
		count(i, i + 1)
	}

	return res
}