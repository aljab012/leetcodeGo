package main

func containsNearbyDuplicate(nums []int, k int) bool {
	l, r := 0, 0
	windowMap := map[int]bool{}
	for r < len(nums) {
		if windowMap[nums[r]] {
			return true
		}
		windowMap[nums[r]] = true
		r++
		if r-l > k {
			delete(windowMap, nums[l])
			l++
		}
	}
	return false
}
