package math

func maximum(a int, b int) int {
	dif := a - b
	k := int(uint(dif) >> 63)
	return a*(k^1) + b*k
}
