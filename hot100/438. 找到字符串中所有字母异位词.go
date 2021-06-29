package hot100

func findAnagrams(s string, p string) []int {
	var res []int

	check := func(a string) bool {
		var flag [26]int
		for i := 0; i < len(p); i++ {
			flag[a[i]-'a']++
			flag[p[i]-'a']--
		}

		for i := range flag {
			if flag[i] != 0 {
				return false
			}
		}
		return true
	}

	for i := 0; i <= len(s)-len(p); i++ {
		subStr := s[i : i+len(p)]
		if check(subStr) {
			res = append(res, i)
		}
	}

	return res
}
