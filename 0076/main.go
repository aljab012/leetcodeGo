package main

import "math"

func minWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}
	ret := ""
	left, count, minLength := 0, 0, math.MaxInt
	LetterMap := make([]int, 128)

	for _, c := range t {
		LetterMap[c]++
	}

	for right := 0; right < len(s); right++ {
		c := s[right]
		LetterMap[c]--
		if LetterMap[c] >= 0 {
			count++
		}
		for count == len(t) {
			// update the res
			windowLength := right - left + 1
			if windowLength < minLength {
				minLength = windowLength
				ret = s[left : right+1]
			}
			// remove the left
			LetterMap[s[left]]++
			if LetterMap[s[left]] > 0 {
				count--
			}
			left++
		}
	}
	return ret
}
