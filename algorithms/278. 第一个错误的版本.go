package algorithms

func isBadVersion(n int) bool {
	return true
}

func firstBadVersion(n int) int {
	l, r := 0, n
	res := -1

	for l <= r {
		mid := l + (r-l)/2
		if isBadVersion(mid) {
			res = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return res
}
