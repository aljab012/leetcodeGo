package main

func isBadVersion(version int) bool {
	return true
}
func firstBadVersion(n int) int {
	l, r := 1, n

	for l < r {
		mid := l + (r-l)/2
		if isBadVersion(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}
