package dp

/**
https://leetcode-cn.com/problems/gu-piao-de-zui-da-li-run-lcof/comments/
*/
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxProfit(prices []int) int {
	size := len(prices)
	if size == 0 || size == 1 {
		return 0
	}

	cur := prices[0]
	res := 0
	for i := 1; i < size; i++ {
		res = max(res, prices[i]-cur)
		cur = min(cur, prices[i])
	}

	return res
}
