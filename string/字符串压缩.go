package string

import "strconv"

/**
https://leetcode-cn.com/problems/compress-string-lcci/comments/
 */

func compressString(S string) string {
	size := len(S)
	var zip string

	for i := 0; i < size; {
		zip += string(S[i])
		idx := i + 1
		if idx == size {
			zip += "1"
			break
		}
		for idx < size && S[idx] == S[i] {
			idx++
		}
		zip += strconv.Itoa(idx - i)
		i = idx
	}
	if len(zip) < len(S) {
		return zip
	}

	return S
}