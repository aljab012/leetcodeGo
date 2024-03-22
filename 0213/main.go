package main

/*
 * Top-down approach recursive
 */
func rob1(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	return max(robOneR(nums[:len(nums)-1]), robOneR(nums[1:]))
}

func robOneR(nums []int) int {
	var fn func(i int) int
	fn = func(i int) int {
		if i >= len(nums) {
			return 0
		}
		return max(nums[i]+fn(i+2), fn(i+1))
	}
	return fn(0)
}

/*
 * Top-down approach recursive with memoization
 */

func rob2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	return max(robOneM(nums[:len(nums)-1]), robOneM(nums[1:]))
}

func robOneM(nums []int) int {
	memo := map[int]int{}
	var fn func(i int) int
	fn = func(i int) int {
		if val, ok := memo[i]; ok {
			return val
		}
		if i >= len(nums) {
			return 0
		}
		ret := max(nums[i]+fn(i+2), fn(i+1))
		memo[i] = ret
		return ret
	}
	return fn(0)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
