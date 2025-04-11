package main

/*
 * Recursive solution
 */
func findTargetSumWays1(nums []int, target int) int {
	var fn func(sum int, i int) int
	fn = func(sum int, i int) int {
		if i == len(nums) && sum == target {
			return 1
		}
		if i >= len(nums) {
			return 0
		}
		return fn(sum+nums[i], i+1) + fn(sum+(-1*nums[i]), i+1)
	}
	return fn(0, 0)
}

/*
 * Recursive with memo
 */
func findTargetSumWays2(nums []int, target int) int {
	type key struct{ sum, i int }
	memo := map[key]int{}
	var fn func(sum int, i int) int
	fn = func(sum int, i int) int {
		if v, ok := memo[key{sum, i}]; ok {
			return v
		}
		if i == len(nums) && sum == target {
			return 1
		}
		if i >= len(nums) {
			return 0
		}
		ret := fn(sum+nums[i], i+1) + fn(sum+(-1*nums[i]), i+1)
		memo[key{sum, i}] = ret
		return ret
	}
	return fn(0, 0)
}
