package hot100

func findDuplicate(nums []int) int {
	slow, fast := 0, 0
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]

		if slow == fast {
			fast = 0
			for nums[slow] != nums[fast] {
				slow = nums[slow]
				fast = nums[fast]
			}
			return nums[slow]
		}
	}
	return -1
}
