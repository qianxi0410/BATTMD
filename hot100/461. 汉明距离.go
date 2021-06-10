package hot100

func hammingDistance(x int, y int) int {
	target, res := x ^ y, 0

	for target > 0 {
		res += target & 0x1
		target >>= 1
	}

	return res
}
