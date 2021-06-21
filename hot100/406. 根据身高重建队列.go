package hot100

import "sort"

func reconstructQueue(people [][]int) [][]int {
	res := make([][]int, len(people))

	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] > people[j][1]
		} else {
			return people[i][0] < people[i][0]
		}
	})

	for _, person := range people {
		spaces := person[1] + 1
		for i := range res {
			if res[i] == nil {
				spaces--
				if spaces == 0 {
					res[i] = person
					break
				}
			}
		}
	}
	return res
}
