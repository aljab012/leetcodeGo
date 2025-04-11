package main

import "sort"

func longestWord(words []string) string {
	sort.Strings(words)
	ret := ""
	set := make(map[string]bool)

	for _, word := range words {
		if len(word) == 1 || set[word[:len(word)-1]] {
			set[word] = true
			if len(word) > len(ret) {
				ret = word
			}
		}
	}

	return ret
}
