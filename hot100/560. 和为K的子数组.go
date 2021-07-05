package hot100

func subarraySum(nums []int, k int) int {
	res := 0
	S := make([]int, len(nums)+1)

	for i := 1; i <= len(nums); i++ {
		S[i] = S[i-1] + nums[i-1]
	}

	for i := 0; i < len(S); i++ {
		for j := i + 1; j < len(S); j++ {
			gap := S[j] - S[i]
			if gap == k {
				res++
			}
		}
	}
	return res
}
