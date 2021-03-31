package string

import (
	"strings"
)

func reverse(s string) string {
	res := ""
	for i := len(s) - 1; i >= 0; i-- {
		res += string(s[i])
	}
	return res
}

func reverseWords(s string) string {
	strs := strings.Split(s, " ")
	for i, str := range strs {
		strs[i] = reverse(str)
	}
	res := strings.Join(strs, " ")
	return res
}
