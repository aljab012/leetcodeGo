package main

/*
 * using a map to store the number and check if the number is already in the map
 * Space complexity: O(n)
 * Time complexity: O(n)
 */
func containsDuplicate1(nums []int) bool {
	set := map[int]bool{}
	for _, n := range nums {
		if _, ok := set[n]; ok {
			return true
		}
		set[n] = true
	}
	return false
}

/*
 * Brute force solution with two loops
 * Space complexity: O(1)
 * Time complexity: O(n^2)
 */
func containsDuplicate2(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}
