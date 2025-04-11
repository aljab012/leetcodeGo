package main

import "math"

func maxNumberOfBalloons(text string) int {
	cMap := make(map[rune]int)
	for _, c := range text {
		cMap[c]++
	}
	return MinValue(cMap['b'], cMap['a'], cMap['l']/2, cMap['o']/2, cMap['n'])
}
func MinValue(values ...int) int {
	min := math.MaxInt
	for _, n := range values {
		if n < min {
			min = n
		}
	}
	return min
}
