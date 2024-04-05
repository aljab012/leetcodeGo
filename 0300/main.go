package main

import "math"

/*
 * Top-down approach
 * Time complexity: O(2^n)
 * Space complexity: O(n)
 * where n is the length of the input array
 * We have 2 choices for each element, either pick it or skip it
 */
func lengthOfLIS1(nums []int) int {
	// brute force top-down
	// either pick the current number or skip it
	var fn func(i int, prev int) int
	fn = func(i int, prev int) int {
		// base cases
		if i == len(nums) {
			return 0
		}
		// pick
		pick, skip := 0, 0
		if nums[i] > prev {
			pick = 1 + fn(i+1, nums[i])
		}
		// skip
		skip = fn(i+1, prev)
		return max(pick, skip)
	}
	return fn(0, math.MinInt)
}

/*
 * Top-down approach with memoization
 */
func lengthOfLIS2(nums []int) int {
	type key struct{ i, prev int }
	memo := map[key]int{}
	var fn func(i int, prev int) int
	fn = func(i int, prev int) int {
		if val, ok := memo[key{i, prev}]; ok {
			return val
		}
		// base cases
		if i == len(nums) {
			return 0
		}
		// pick
		pick, skip := 0, 0
		if nums[i] > prev {
			pick = 1 + fn(i+1, nums[i])
		}
		// skip
		skip = fn(i+1, prev)
		ret := max(pick, skip)
		memo[key{i, prev}] = ret
		return ret
	}
	return fn(0, math.MinInt)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
