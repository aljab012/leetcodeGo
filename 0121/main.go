package main

import "math"

func maxProfit(prices []int) int {

	minPrice := math.MaxInt
	var maxProfit int

	for _, p := range prices {
		minPrice = min(minPrice, p)
		maxProfit = max(maxProfit, p-minPrice)

	}
	return maxProfit
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
