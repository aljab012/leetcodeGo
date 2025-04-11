package main

func change1(amount int, coins []int) int {
	var fn func(int, int) int
	fn = func(rem int, i int) int {
		// base cases
		if i >= len(coins) {
			return 0
		}
		if rem == 0 {
			return 1
		}

		// recursive
		ret := 0
		for ; i < len(coins); i++ {
			if rem-coins[i] >= 0 {
				ret += fn(rem-coins[i], i)
			}
		}
		return ret
	}
	return fn(amount, 0)
}
func change(amount int, coins []int) int {
	type key struct {
		rem int
		i   int
	}
	memo := map[key]int{}
	var fn func(int, int) int
	fn = func(rem int, i int) int {
		if val, ok := memo[key{rem, i}]; ok {
			return val
		}
		// base cases
		if i >= len(coins) {
			return 0
		}
		if rem == 0 {
			return 1
		}

		// recursive
		ret := 0
		for ; i < len(coins); i++ {
			if rem-coins[i] >= 0 {
				ret += fn(rem-coins[i], i)
			}
		}
		memo[key{rem, i}] = ret
		return ret
	}
	return fn(amount, 0)
}
