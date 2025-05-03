package main

func maximumSubarraySum(nums []int, k int) int64 {
	result := 0

	l, r := 0, 0
	windowMap := map[int]int{}
	windowSum := 0
	for r < len(nums) {
		// grow window
		windowMap[nums[r]]++
		windowSum += nums[r]
		r++
		if r-l == k {
			// shrink window
			if len(windowMap) == k {
				result = max(result, windowSum)
			}
			windowMap[nums[l]]--
			if windowMap[nums[l]] == 0 {
				delete(windowMap, nums[l])
			}
			windowSum -= nums[l]
			l++
		}
	}
	return int64(result)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
