package main

func maxProfit1(prices []int) int {
	var fn func(int, bool) int
	fn = func(i int, canBuy bool) int {
		if i >= len(prices) {
			return 0
		}

		if canBuy { // only buy or cooldown
			return max(fn(i+1, false)-prices[i], fn(i+1, canBuy))
		} else { // only sell or cooldown
			return max(fn(i+2, true)+prices[i], fn(i+1, canBuy))
		}
	}
	return fn(0, true)
}
func maxProfit2(prices []int) int {
	type key struct {
		i      int
		canBuy bool
	}
	memo := map[key]int{}
	var fn func(int, bool) int
	fn = func(i int, canBuy bool) int {
		if val, ok := memo[key{i, canBuy}]; ok {
			return val
		}
		if i >= len(prices) {
			return 0
		}
		ret := 0
		if canBuy { // only buy or cooldown
			ret = max(fn(i+1, false)-prices[i], fn(i+1, canBuy))
		} else { // only sell or cooldown
			ret = max(fn(i+2, true)+prices[i], fn(i+1, canBuy))
		}
		memo[key{i, canBuy}] = ret
		return ret
	}
	return fn(0, true)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
