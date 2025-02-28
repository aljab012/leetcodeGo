package main

func removeElement(nums []int, val int) int {
	pivot := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			continue
		}
		nums[pivot] = nums[i]
		pivot++
	}
	return pivot
}
