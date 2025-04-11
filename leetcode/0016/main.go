package main

import "sort"

func threeSumClosest(nums []int, target int) int {

	sort.Ints(nums)

	closestSum := nums[0] + nums[1] + nums[len(nums)-1]

	for i := 0; i < len(nums); i++ {
		left, right := i+1, len(nums)-1
		for left < right {
			curSum := nums[i] + nums[left] + nums[right]

			if Abs(curSum-target) < Abs(closestSum-target) {
				closestSum = curSum
			}
			if curSum < target {
				left++
			} else {
				right--
			}
		}
	}
	return closestSum
}

func Abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}
