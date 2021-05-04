package hot100

import "sort"

type Intervals [][]int

func (s Intervals) Len() int {
	return len(s)
}
func (s Intervals) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s Intervals) Less(i, j int) bool {
	return s[i][0] < s[j][0]
}

func merge(intervals [][]int) [][]int {
	sort.Sort(Intervals(intervals))

	res := make([][]int, 0, len(intervals))
	for i := 0; i < len(intervals); i++ {
		tmp := make([]int, 2)
		tmp[0] = intervals[i][0]
		tmp[1] = intervals[i][1]
		j := i
		for ; j < len(intervals) && j + 1 < len(intervals); j++ {
			if tmp[1] < intervals[j + 1][0] {
				break
			}
			tmp[1] = max(tmp[1], intervals[j + 1][1])
		}
		i = j
		res = append(res, tmp)
	}
	return res
}
