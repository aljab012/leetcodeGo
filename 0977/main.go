package main

func sortedSquares(nums []int) []int {
	ret := make([]int, len(nums))
	l, r := 0, len(nums)-1

	index := len(nums) - 1
	for index >= 0 {
		lv := nums[l] * nums[l]
		rv := nums[r] * nums[r]
		if lv > rv {
			ret[index] = lv
			l++
			index--
		} else {
			ret[index] = rv
			r--
			index--
		}
	}
	return ret
}
