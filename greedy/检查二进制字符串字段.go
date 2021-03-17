package greedy

func checkOnesSegment(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			for i < len(s) && s[i] == '1' {
				i++
			}
			for j := i + 1; j < len(s); j++ {
				if s[j] == '1' {
					return false
				}
			}
		}
	}
	return true
}
