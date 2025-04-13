package main

func guess(num int) int {
	return 0
}

func guessNumber(n int) int {
	left, right := 1, n
	for left <= right {
		mid := left + (right-left)/2
		result := guess(mid)
		if result == 0 {
			return mid
		} else if result < 0 {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
