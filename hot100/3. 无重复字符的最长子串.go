package hot100

func lengthOfLongestSubstring(s string) int {
	var last [128]int
	for i := 0; i < len(last); i++ {
		last[i] = -1
	}
	var left = -1
	var ans = 0
	for i := 0; i < len(s); i++ {
		var si = int(s[i])
		left = max(left, last[si])
		last[si] = i
		ans = max(ans, i - left)
	}
	return ans
}