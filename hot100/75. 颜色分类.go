package hot100

func sortColors(nums []int)  {
	l, r := 0, len(nums) - 1

	for i := l; i <= r; {
		if nums[i] == 2 {
			nums[i], nums[r] = nums[r], nums[i]
			r--
		} else if nums[i] == 0 {
			nums[l], nums[i] = nums[i], nums[l]
			l++
			i++
		} else {
			i++
		}
	}
}
