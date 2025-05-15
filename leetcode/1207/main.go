package main

func uniqueOccurrences(arr []int) bool {
	countMap := map[int]int{}
	for _, num := range arr {
		countMap[num]++
	}

	seen := map[int]bool{}
	for _, count := range countMap {
		if seen[count] {
			return false
		}
		seen[count] = true
	}
	return true
}
