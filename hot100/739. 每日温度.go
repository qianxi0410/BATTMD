package hot100

func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))

	for i := len(temperatures) - 2; i >= 0; i-- {
		for j := i + 1; j < len(temperatures); j += res[j] {
			if temperatures[i] < temperatures[j] {
				res[i] = j - i
				break
			} else if res[j] == 0 {
				res[i] = 0
				break
			}
		}
	}
	return res
}
