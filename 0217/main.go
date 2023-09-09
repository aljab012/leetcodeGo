package main

// containsDuplicate checks if the given slice of integers contains duplicates.
// It returns true if it does, false otherwise.
// It uses a map to keep track of the numbers it has seen.
// Time complexity: O(n)
// Space complexity: O(n)
func containsDuplicate(nums []int) bool {
	nMap := make(map[int]bool)
	for _, n := range nums {
		_, found := nMap[n]
		if found {
			return true
		}
		nMap[n] = true
	}
	return false
}
