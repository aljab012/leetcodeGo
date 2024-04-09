package main

/*
 * Use two pointers technique
 * Time complexity: O(n)
 * Space complexity: O(1)
 */
func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1

	for left < right {
		sum := numbers[left] + numbers[right]
		if sum < target {
			left++
		} else if sum > target {
			right--
		} else {
			return []int{left + 1, right + 1}
		}
	}
	return []int{-1, -1}
}
