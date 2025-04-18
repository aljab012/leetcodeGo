package main

func isBadVersion(version int) bool {
	return true
}

// good good/ bad bad
func firstBadVersion(n int) int {
	isBefore := func(num int) bool {
		return !isBadVersion(num)
	}

	l, r := 1, n

	// edge cases
	// empty (covered by contraint)
	// the whole range is 'after'
	if isBadVersion(l) {
		return l
	}
	// the whole range is 'before'
	if !isBadVersion(r) {
		return -1
	}

	for (r - l) > 1 {
		mid := (l + r) / 2
		if isBefore(mid) {
			l = mid
		} else {
			r = mid
		}
	}

	return r
}
