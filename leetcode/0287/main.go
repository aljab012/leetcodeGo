package main

/*
 * intuition: the array is a linked list, and the duplicate number is the cycle start of the linked list.
 * 1. find the intersection point of the two runners.
 * 2. find the entrance of the cycle.
 * 3. return the entrance.
 * time complexity: O(n)
 * space complexity: O(1)
 */
func findDuplicate(nums []int) int {
	slow := nums[0]
	fast := nums[0]
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}
	slow = nums[0]
	for fast != slow {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}
