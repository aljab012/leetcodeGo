package main

import "math"

func mincostTickets1(days []int, costs []int) int {
	var fn func(startIndex int) int
	fn = func(startIndex int) int {
		if startIndex == len(days) {
			return 0
		}
		ret := math.MaxInt64
		passes := []int{1, 7, 30}
		for i, p := range passes {
			cost := costs[i]
			passEnd := days[startIndex] + p
			newIndex := startIndex
			for newIndex < len(days) && days[newIndex] < passEnd {
				newIndex++
			}
			ret = min(ret, cost+fn(newIndex))
		}
		return ret
	}
	return fn(0)
}

func mincostTickets2(days []int, costs []int) int {
	memo := map[int]int{}

	var fn func(startIndex int) int
	fn = func(startIndex int) int {
		if v, ok := memo[startIndex]; ok {
			return v
		}
		if startIndex == len(days) {
			return 0
		}
		ret := math.MaxInt64
		passes := []int{1, 7, 30}
		for i, p := range passes {
			passEnd := days[startIndex] + p
			newIndex := startIndex
			for newIndex < len(days) && days[newIndex] < passEnd {
				newIndex++
			}
			ret = min(ret, costs[i]+fn(newIndex))
		}
		memo[startIndex] = ret
		return ret
	}
	return fn(0)
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
