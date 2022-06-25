package main

func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}
	pivot := 2
	for i := 2; i < len(nums); i++ {
		if nums[i] != nums[pivot-2] {
			nums[pivot] = nums[i]
			pivot++
		}
	}
	return pivot
}
