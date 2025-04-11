package main

import (
	"fmt"
	"sort"
)

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

func main() {
	intervals := [][]int{
		{1, 3},
		{2, 6},
		{8, 10},
		{15, 18},
	}
	// expected: [[1 6] [8 10] [15 18]]
	fmt.Println(merge(intervals))
}
