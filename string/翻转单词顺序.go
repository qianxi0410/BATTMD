package string

import (
	"regexp"
	"strings"
)

/**
https://leetcode-cn.com/problems/fan-zhuan-dan-ci-shun-xu-lcof/comments/
 */

func reverseWordsI(s string) string {
	s = strings.Trim(s, " ")
	reg := regexp.MustCompile(`\s+`)
	s = reg.ReplaceAllLiteralString(s, " ")
	strs := strings.Split(s, " ")
	for i := 0; i < len(strs) / 2; i++ {
		strs[i], strs[len(strs) - i - 1] = strs[len(strs) - i - 1], strs[i]
	}
	s = strings.Join(strs, " ")
	return s
}
