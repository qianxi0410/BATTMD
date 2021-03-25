package dp

func numberOf2sInRange(n int) int {
	res := 0
	for base := 1; base <= n; base *= 10 {
		leftPart, rightPart := n / base, n % base
		res += leftPart / 10 * base
		if lastBit := leftPart % 10; lastBit > 2 {
			res += base
		} else if lastBit == 2 {
			res += rightPart + 1
		}
	}
	return res
}
