package main

import "math"

func maxProfit(prices []int) int {
	maxProfit := 0
	minPrice := math.MaxInt

	for _, p := range prices {
		minPrice = min(p, minPrice)
		profit := p - minPrice
		if profit > maxProfit {
			maxProfit = profit
		}
	}
	return maxProfit
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
