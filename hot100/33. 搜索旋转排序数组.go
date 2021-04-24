package hot100

func search(nums []int, target int) int {
	k := 0

	for i := 0; i < len(nums); i++ {
		if i < len(nums) - 1 && nums[i] > nums[i + 1] {
			k = i
			break
		}
	}
	l, r := -1, -1
	if target < nums[0] {
		l = k + 1
		r = len(nums) - 1
	} else {
		l = 0
		r = k
	}

	for l < r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}
