package hot100


func canFinish(numCourses int, prerequisites [][]int) bool {
	T := make([][]int, numCourses)
	inDegree := make([]int, numCourses)
	res := make([]int, 0, numCourses)

	for _, v := range prerequisites {
		T[v[1]] = append(T[v[1]], v[0])
		inDegree[v[0]]++
	}

	var q []int
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			q = append(q, i)
		}
	}

	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		res = append(res, u)
		for _, v := range T[u] {
			inDegree[v]--
			if inDegree[v] == 0 {
				q = append(q, v)
			}
		}
	}
	return len(res) == numCourses
}
