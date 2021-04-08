package binary

func search(nums []int, target int) int {
	n := len(nums)
	if n <= 0 {
		return 0
	}

	left, right, mid, times := 0, n-1, 0, 0

	for left <= right {
		mid = (left + right) / 2
		if nums[mid] == target {
			times++
			break
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	// 没有找到目标值
	if left > right {
		return 0
	}

	// 向左查找
	for i := mid-1; i >= 0; i-- {
		if nums[i] == target {
			times++
		} else {
			break
		}
	}

	// 向右查找
	for i := mid + 1; i < n; i++ {
		if nums[i] == target {
			times++
		} else {
			break
		}
	}

	return times
}
