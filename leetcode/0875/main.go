package main

import "math"

func minEatingSpeed(piles []int, h int) int {
	canFinish := func(k int) bool {
		totalHours := 0
		for _, p := range piles {
			totalHours += int(math.Ceil(float64(p) / float64(k)))
		}
		return totalHours <= h
	}

	maxPile := 0
	for _, pile := range piles {
		if pile > maxPile {
			maxPile = pile
		}
	}

	l, r := 1, maxPile
	for l < r {
		mid := (l + r) / 2
		if canFinish(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}
