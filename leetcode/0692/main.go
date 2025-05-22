package main

import "sort"

func topKFrequent(words []string, k int) []string {
	frequency := map[string]int{}
	for _, word := range words {
		frequency[word]++
	}
	uniqueWords := []string{}
	for key := range frequency {
		uniqueWords = append(uniqueWords, key)
	}
	sort.Slice(uniqueWords, func(i, j int) bool {
		if frequency[uniqueWords[i]] == frequency[uniqueWords[j]] {
			return uniqueWords[i] < uniqueWords[j]
		}
		return frequency[uniqueWords[i]] > frequency[uniqueWords[j]]
	})
	return uniqueWords[:k]
}
