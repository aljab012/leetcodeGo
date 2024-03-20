package main

import "math"

/*
 * Top-down approach recursive
 */
func coinChange1(coins []int, amount int) int {
	var fn func(int) int
	fn = func(rem int) int {
		if rem == 0 {
			return 0
		}
		ret := math.MaxInt - 1
		for _, c := range coins {
			if c > rem {
				continue
			}
			ret = min(ret, 1+fn(rem-c))
		}
		return ret
	}
	ret := fn(amount)
	if ret > amount {
		return -1
	}
	return ret
}

/*
 * Top-down approach recursive with memoization
 */
func coinChange2(coins []int, amount int) int {
	memo := map[int]int{}
	var fn func(int) int
	fn = func(rem int) int {
		if val, ok := memo[rem]; ok {
			return val
		}
		if rem == 0 {
			return 0
		}
		ret := math.MaxInt - 1
		for _, c := range coins {
			if c > rem {
				continue
			}
			ret = min(ret, 1+fn(rem-c))
		}
		memo[rem] = ret
		return ret
	}
	ret := fn(amount)
	if ret > amount {
		return -1
	}
	return ret
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
