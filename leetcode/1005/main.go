package main

import (
	"math"
	"sort"
)

// idea: flip negatives
// then flip the min val once if k is odd
func largestSumAfterKNegations(nums []int, k int) int {
	sort.Ints(nums)
	for i := 0; i < len(nums) && k > 0 && nums[i] < 0; i++ {
		nums[i] = -nums[i]
		k--
	}
	ret, min := 0, math.MaxInt
	for i := range nums {
		ret += nums[i]
		min = Min(min, nums[i])
	}
	return ret - (k%2)*min*2
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
