package hot100

func longestPalindrome(s string) string {
	low, maxLen := 0, 0
	for i := range s {
		expand(s, i, i, &low, &maxLen)
		expand(s, i, i+1, &low, &maxLen)
	}
	return s[low : low+maxLen]
}
func expand(s string, l, r int, low, maxLen *int) {
	for ; l >= 0 && r < len(s) && s[l] == s[r]; l, r = l-1, r+1 {
	}
	if *maxLen < r-l-1 {
		*low = l + 1
		*maxLen = r - l - 1
	}
}
