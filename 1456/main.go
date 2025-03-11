package main

func maxVowels(s string, k int) int {
	vowels := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}

	maxCount := 0
	count := 0

	l, r := 0, 0
	for r < len(s) {
		// grow window
		if _, ok := vowels[s[r]]; ok {
			count++
		}
		r++
		if r-l == k {
			maxCount = max(maxCount, count)
			// shrink window
			if _, ok := vowels[s[l]]; ok {
				count--
			}
			l++
		}

	}
	return maxCount
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
