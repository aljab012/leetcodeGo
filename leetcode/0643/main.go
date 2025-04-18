package main

import "math"

func findMaxAverage(nums []int, k int) float64 {
	l, r := 0, 0
	maxSum := math.MinInt
	curSum := 0
	for r < len(nums) {
		curSum += nums[r]
		r++
		if r-l == k {
			maxSum = max(maxSum, curSum)
			curSum -= nums[l]
			l++
		}
	}
	return float64(maxSum) / float64(k)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
