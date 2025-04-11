package main

/*
 * solution using two loops
 * Space complexity: O(1)
 * Time complexity: O(n^2)
 */
func twoSum1(nums []int, target int) []int {
	for i, n := range nums {
		for j := i + 1; j < len(nums); j++ {
			if n+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{-1, -1}
}

/*
 * solution using a map to store all the numbers and check if the complement is in the map
 * Space complexity: O(n)
 * Time complexity: O(n)
 */
func twoSum2(nums []int, target int) []int {
	nMap := map[int]int{}
	for i, n := range nums {
		nMap[n] = i
	}
	for i, n := range nums {
		comp := target - n
		if val, ok := nMap[comp]; ok && val != i {
			return []int{i, val}
		}
	}
	return []int{-1, -1}
}

/*
 * using a map to store the number and check if the number is already in the map
 * Space complexity: O(n)
 * Time complexity: O(n)
 */
// Idea: interate through the array, see if you have encountered a complement in the previous nums.
func twoSum3(nums []int, target int) []int {
	nMap := map[int]int{}
	for i, n := range nums {
		comp := target - n
		if val, ok := nMap[comp]; ok {
			return []int{val, i}
		}
		nMap[n] = i
	}
	return []int{-1, -1}
}
