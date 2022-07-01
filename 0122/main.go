package main

// idea: Draw the timeline to solve it
func maxProfit(prices []int) int {
	ret := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			ret += prices[i] - prices[i-1]
		}
	}
	return ret
}
