package main

func mergeAlternately(word1 string, word2 string) string {
	ret := ""

	p1, p2 := 0, 0
	isFirst := true
	for p1 < len(word1) || p2 < len(word2) {
		if isFirst && p1 < len(word1) {
			ret = ret + string(word1[p1])
			p1++
		}
		if !isFirst && p2 < len(word2) {
			ret = ret + string(word2[p2])
			p2++
		}
		isFirst = !isFirst
	}
	return ret
}
