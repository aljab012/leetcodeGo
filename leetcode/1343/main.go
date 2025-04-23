package main

func numOfSubarrays(arr []int, k int, threshold int) int {
	l, r := 0, 0
	windowSum := 0
	count := 0
	for r < len(arr) {
		windowSum += arr[r]
		r++
		if r-l == k {
			windowThreshold := float64(windowSum) / float64(k)
			if windowThreshold >= float64(threshold) {
				count++
			}
			windowSum -= arr[l]
			l++
		}
	}
	return count
}
