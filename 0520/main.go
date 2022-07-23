package main

import "unicode"

func detectCapitalUse(word string) bool {
	count := 0

	for _, c := range word {
		if unicode.IsUpper(c) {
			count++
		}
	}
	// valid: all upper or all lower
	if count == len(word) || count == 0 {
		return true
	}
	// valid: if one char is upper, it must be the first
	if count == 1 && unicode.IsUpper(rune(word[0])) {
		return true
	}
	return false
}
