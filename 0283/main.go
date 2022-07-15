package main

func moveZeroes(nums []int) {
	index := 0

	for _, n := range nums {
		if n != 0 {
			nums[index] = n
			index++
		}
	}
	for index < len(nums) {
		nums[index] = 0
		index++
	}
}
