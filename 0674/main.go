package main

func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	curMax, max := 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			curMax++
			max = Max(max, curMax)
		} else {
			curMax = 1
		}
	}
	return max
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
