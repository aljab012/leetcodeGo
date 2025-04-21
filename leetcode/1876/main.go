package main

func countGoodSubstrings(s string) int {
	l, r := 0, 0
	windowMap := map[byte]int{}
	countStrings := 0

	for r < len(s) {
		windowMap[s[r]]++
		r++
		if r-l == 3 {
			if len(windowMap) == 3 {
				countStrings++
			}
			windowMap[s[l]]--
			if windowMap[s[l]] == 0 {
				delete(windowMap, s[l])
			}
			l++
		}
	}
	return countStrings
}
