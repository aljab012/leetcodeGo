package main

func lengthOfLongestSubstring(s string) int {
	curBest := 0
	set := map[byte]bool{}

	l, r := 0, 0
	for r < len(s) {
		if _, ok := set[s[r]]; !ok {
			set[s[r]] = true
			r++
			curBest = max(curBest, r-l)
		} else {
			delete(set, s[l])
			l++
		}
	}
	return curBest
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
