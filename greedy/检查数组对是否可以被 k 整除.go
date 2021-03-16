package greedy

func canArrange(arr []int, k int) bool {
	m := map[int]int{}
	for _, a := range arr {
		m[(a%k+k)%k]++
	}
	for key, value := range m {
		if key == 0 {
			if value%2 == 1 {
				return false
			}
			continue
		}
		if m[k-key] != value {
			return false
		}
	}
	return true
}
