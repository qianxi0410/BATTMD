package hot100

import "strings"

func wordBreak(s string, wordDict []string) bool {
	n, memo := len(s), make([]bool, len(s)+1)

	contain := func(s string) bool {
		for i := range wordDict {
			if strings.Compare(s, wordDict[i]) == 0 {
				return true
			}
		}
		return false
	}

	memo[0] = true
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			if memo[j] && contain(s[j:i]) {
				memo[i] = true
				break
			}
		}
	}
	return memo[n]
}
