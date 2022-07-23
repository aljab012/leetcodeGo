package main

func firstUniqChar(s string) int {
	cMap := make([]int, 28)

	for _, c := range s {
		cMap[int(c-'a')]++
	}
	for i, c := range s {
		if cMap[int(c-'a')] == 1 {
			return i
		}
	}
	return -1
}
