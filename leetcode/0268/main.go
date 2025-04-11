package main

// idea: add values from 0 to n then subtract nums[i] values to find the missing
func missingNumber(nums []int) int {
	ret := len(nums)
	for i := 0; i < len(nums); i++ {
		ret += (i - nums[i])
	}
	return ret
}
