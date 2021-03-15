package greedy

/**
https://leetcode-cn.com/problems/can-place-flowers/comments/
*/
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	res := 0
	for i := range flowerbed {
		if flowerbed[i] == 0 {
			if flowerbed[max(i-1, 0)] == 0 && flowerbed[min(i+1, len(flowerbed)-1)] == 0 {
				flowerbed[i] = 1
				res++
			}
		} else {
			res++
		}
	}
	return res >= n
}
