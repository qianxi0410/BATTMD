package hot100

func _maxProfit(prices []int) int {
	size := len(prices)
	sell, buy, freeze := make([]int, size), make([]int, size), make([]int, size)

	if size == 0 {
		return 0
	}
	buy[0] = -prices[0]

	for i := 1; i < size; i++ {
		sell[i] = buy[i - 1] + prices[i]
		buy[i] = max(freeze[i - 1] - prices[i], buy[i - 1])
		freeze[i] = max(sell[i - 1], freeze[i - 1])
	}

	return max(sell[size - 1], freeze[size - 1])
}
