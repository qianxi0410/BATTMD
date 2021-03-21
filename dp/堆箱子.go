package dp

import "sort"

/**
https://leetcode-cn.com/problems/pile-box-lcci/submissions/
*/
func pileBox(box [][]int) int {
	boxLen := len(box)
	if boxLen == 0 {
		return 0
	}
	sort.Slice(box, func(i, j int) bool {
		return box[i][0] < box[j][0]
	})
	dp, res := make([]int, boxLen), box[0][2]
	dp[0] = box[0][2]
	for i := 1; i < boxLen; i++ {
		dp[i] = box[i][2]
		for j := 0; j < i; j++ {
			if box[i][0] > box[j][0] && box[i][1] > box[j][1] && box[i][2] > box[j][2] {
				dp[i] = max(dp[i], dp[j]+box[i][2])
			}
		}
		res = max(res, dp[i])
	}
	return res
}
