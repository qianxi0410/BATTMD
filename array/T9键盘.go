package main

/**
https://leetcode-cn.com/problems/t9-lcci/comments/
*/

func getValidT9Words(num string, words []string) []string {
	res := make([]string, 0, len(words))

	for _, v := range words {
		tmp := make([]byte, 0, len(num))
		for _, v := range v {
			if v <= 'r' {
				gap := (v-'a')/3 + 2
				tmp = append(tmp, byte(gap)+'0')
			} else if v == 's' {
				tmp = append(tmp, '7')
			} else if v >= 't' && v <= 'v' {
				tmp = append(tmp, '8')
			} else {
				tmp = append(tmp, '9')
			}

		}
		if string(tmp) == num {
			res = append(res, v)
		}
	}
	return res

}
