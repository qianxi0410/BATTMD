package string

import "strings"

/**
https://leetcode-cn.com/problems/string-rotation-lcci/submissions/
 */

func isFlipedString(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	s := s2 + s2
	return strings.Contains(s, s1)
}
