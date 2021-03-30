package string
/**
https://leetcode-cn.com/problems/check-permutation-lcci/comments/
 */

func CheckPermutation(s1 string, s2 string) bool {
	arr := make([]int, 126)

	if len(s1) != len(s2) {
		return false
	}

	for i := 0; i < len(s1); i++ {
		arr[s1[i]]++
		arr[s2[i]]--
	}

	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			return false
		}
	}
	return true
}
