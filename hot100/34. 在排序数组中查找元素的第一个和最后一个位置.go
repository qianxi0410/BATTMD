package hot100

func searchRange(nums []int, target int) []int {
	l, r := 0, len(nums) - 1
	res := []int{-1, -1}

	for l <= r {
		mid := l + (r - l) >> 1
		if nums[mid] == target {
			j := mid + 1
			for j <= r && nums[j] == target {
				j++
			}
			res[1] = j - 1
			j = mid - 1
			for j >= l && nums[j] == target {
				j--
			}
			res[0] = j + 1
			return res
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return res
}
