package main

import "math"

func shortestToChar(s string, c byte) []int {
	ret := make([]int, len(s))

	pos := math.MinInt / 2

	for i := range ret {
		if s[i] == c {
			pos = i
		}
		ret[i] = i - pos
	}
	pos = math.MaxInt / 2
	for i := len(ret) - 1; i >= 0; i-- {
		if s[i] == c {
			pos = i
		}
		ret[i] = Min(ret[i], pos-i)
	}
	return ret
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
