package main

/**
https://leetcode-cn.com/problems/living-people-lcci/
*/

func maxAliveYear(birth []int, death []int) int {
	cnt := make([]int, 102)
	tmp, max, idx := 0, 0, 0

	for i := 0; i < len(birth); i++ {
		cnt[birth[i]-1900]++
		cnt[death[i]-1900+1]--
	}

	for i := 0; i < 101; i++ {
		tmp += cnt[i]
		if tmp > max {
			max = tmp
			idx = i
		}
	}
	return idx + 1900
}
