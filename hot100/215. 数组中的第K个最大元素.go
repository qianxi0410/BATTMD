package hot100


func findKthLargest(nums []int, k int) int {
	var helper func(arr []int, l, r int) int

	helper = func(arr []int, l, r int) int {
		tmp := nums[l]

		for l < r {
			for nums[r] <= tmp && r > l {
				r--
			}
			nums[r], nums[l] = nums[l], nums[r]

			for nums[l] >= tmp && r > l {
				l++
			}
			nums[r], nums[l] = nums[l], nums[r]
		}
		nums[l] = tmp
		return l
	}

	var foo func(l, r int)
	foo = func(l, r int) {
		idx := helper(nums, l, r)
		if idx == k - 1 {
			return
		} else if idx > k - 1 {
			foo(l, idx - 1)
		} else {
			foo(idx + 1, r)
		}
	}

	foo(0, len(nums) - 1)

	return nums[k - 1]
}