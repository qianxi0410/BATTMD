package hot100

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	var f func(string, int, int)

	f = func(tmp string, l int, r int) {
		if l < r || l > n || r > n {
			return
		}
		if l == r && l == n {
			res = append(res, tmp)
		}
		f(tmp + "(", l + 1, r)
		f(tmp + ")", l , r + 1)
	}
	f("", 0, 0)
	return res
}
