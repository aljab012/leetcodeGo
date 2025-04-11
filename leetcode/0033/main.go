package main

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		mid := l + (r-l)/2
		if target == nums[mid] {
			return mid
		}
		// we are in the left sorted
		if nums[mid] >= nums[0] {
			if target < nums[mid] && target >= nums[0] { // in the left range
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else { // we are in right sorted
			if target > nums[mid] && target <= nums[len(nums)-1] { // in the right range
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}
