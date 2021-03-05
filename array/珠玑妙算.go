package main

/**
https://leetcode-cn.com/problems/master-mind-lcci/
*/

func masterMind(solution string, guess string) []int {
	size := 4
	m := make(map[rune]int, size)
	m['R'], m['Y'], m['G'], m['B'] = 0, 0, 0, 0

	res := make([]int, 2)

	for _, v := range solution {
		m[v]++
	}

	flag := make([]bool, size)

	for i := 0; i < size; i++ {
		if solution[i] == guess[i] {
			res[0]++
			m[rune(solution[i])]--
			flag[i] = true
		}
	}

	for i, v := range guess {
		if !flag[i] && m[v] > 0 {
			res[1]++
			m[v]--
		}
	}

	return res
}
