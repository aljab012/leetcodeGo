package main

/*
 * Top-down approach recursive
 * Time complexity: O(2^n)
 * Space complexity: O(n)
 * Intuition: At each step, we can either take the current number or skip it.
 */
func canPartition1(nums []int) bool {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	if sum%2 != 0 {
		return false
	}
	target := sum / 2
	var fn func(int, int) bool
	fn = func(i int, lTarget int) bool {
		// base cases
		if i >= len(nums) || lTarget < 0 {
			return false
		}
		if lTarget == 0 {
			return true
		}

		// recursive
		// take i
		take := fn(i+1, lTarget-nums[i])
		// skip i
		skip := fn(i+1, lTarget)

		return take || skip
	}
	return fn(0, target)
}

/*
 * Top-down approach recursive with memoization
 */
func canPartition2(nums []int) bool {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	if sum%2 != 0 {
		return false
	}
	target := sum / 2
	type key struct {
		i, lTarget int
	}
	memo := map[key]bool{}
	var fn func(int, int) bool
	fn = func(i int, lTarget int) bool {
		if val, ok := memo[key{i, lTarget}]; ok {
			return val
		}
		// base cases
		if i >= len(nums) || lTarget < 0 {
			return false
		}
		if lTarget == 0 {
			return true
		}

		// recursive
		// take i
		take := fn(i+1, lTarget-nums[i])
		// skip i
		skip := fn(i+1, lTarget)

		ret := take || skip
		memo[key{i, lTarget}] = ret
		return ret
	}
	return fn(0, target)
}
