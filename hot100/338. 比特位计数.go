package hot100

func countBits(n int) []int {
	res := make([]int, n + 1)
	for i := 1; i <= n; i++ {
		res[i] = res[i >> 1] + (i & 0x1)
	}

	return res
}