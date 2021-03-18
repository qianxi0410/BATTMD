package greedy

/**
https://leetcode-cn.com/problems/water-bottles/comments/
*/
func numWaterBottles(numBottles int, numExchange int) int {
	res := numBottles

	for numBottles >= numExchange {
		tmp := numBottles / numExchange
		res += tmp
		numBottles -= tmp * numExchange
		numBottles += tmp
	}
	return res
}
