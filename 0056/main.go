package main

import "sort"

func merge(intervals [][]int) [][]int {

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	i := 1
	for i < len(intervals) {
		if intervals[i-1][1] >= intervals[i][0] {
			intervals[i-1][1] = Max(intervals[i][1], intervals[i-1][1])
			intervals = append(append([][]int{}, intervals[:i]...), intervals[i+1:]...)
		} else {
			i++
		}
	}

	return intervals
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
