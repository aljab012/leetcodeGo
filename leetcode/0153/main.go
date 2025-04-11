package main

func findMin(nums []int) int {
	if len(nums) == 1 || nums[0] < nums[len(nums)-1] {
		return nums[0]
	}
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] > nums[mid+1] {
			return nums[mid+1]
		} else if nums[mid] < nums[mid-1] {
			return nums[mid]
		}
		if nums[0] < nums[mid] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}
