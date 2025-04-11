package main

/*
 * Binary search
 * Space complexity: O(1)
 * Time complexity: O(log n)
 * The array must be sorted
 * Intuition: We can use binary search to find the target number
 * If the target number is found, return the index
 * If the target number is not found, return -1
 * The time complexity is O(log n) because we are dividing the search space in half each time
 * The space complexity is O(1) because we are not using any extra space
 * The array must be sorted
 * The algorithm works because we are dividing the search space in half each time.
 */
func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if target == nums[mid] {
			return mid
		} else if target > nums[mid] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}
