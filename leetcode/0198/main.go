package main

/*
 * Recursive approach
 */
func rob1(nums []int) int {
	var fn func(i int) int
	fn = func(i int) int {
		// base case - reach the end
		if i >= len(nums) {
			return 0 // no more money
		}
		// recursive call - return the max of robbing and not robbing the current index
		return max(nums[i]+fn(i+2), fn(i+1))
	}
	return fn(0)
}

/*
 * Recursive with memo
 */
func rob2(nums []int) int {
	memo := map[int]int{}
	var fn func(i int) int
	fn = func(i int) int {
		if v, ok := memo[i]; ok {
			return v
		}
		// base case - reach the end
		if i >= len(nums) {
			return 0 // no more money
		}
		// recursive call - return the max of robbing and not robbing the current index
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
