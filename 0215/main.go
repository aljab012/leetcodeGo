package main

func findKthLargest(nums []int, k int) int {
	k = len(nums) - k
	return helper(nums, k, 0, len(nums)-1)
}

func helper(nums []int, k int, left, right int) int {
	pivot, ptr := nums[right], left
	for i := left; i < right; i++ {
		if nums[i] <= pivot {
			nums[ptr], nums[i] = nums[i], nums[ptr]
			ptr++
		}
	}
	nums[ptr], nums[right] = nums[right], nums[ptr]
	if k < ptr {
		return helper(nums, k, left, ptr-1)
	} else if k > ptr {
		return helper(nums, k, ptr+1, right)
	} else {
		return nums[ptr]
	}
}
