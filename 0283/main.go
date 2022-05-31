package main

func moveZeroes(nums []int) {
	index := 0
	for i := range nums {
		if nums[i] == 0 {
			continue
		}
		nums[index] = nums[i]
		index += 1
	}
	for index < len(nums) {
		nums[index] = 0
		index += 1
	}

}
