package hot100

func maxProfit(prices []int) int {
	res := 0
	minVal := prices[0]

	for i := 1; i < len(prices); i++ {
		res = max(prices[i] - minVal, res)
		minVal = min(minVal, prices[i])
	}

	return res
}
