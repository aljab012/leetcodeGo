package main

func peakIndexInMountainArray(arr []int) int {
	isBefore := func(i int) bool {
		return i == 0 || arr[i] > arr[i-1]
	}
	left, right := 0, len(arr)-1
	for (right - left) > 1 {
		mid := (left + right) / 2
		if isBefore(mid) {
			left = mid
		} else {
			right = mid
		}
	}
	return left
}
