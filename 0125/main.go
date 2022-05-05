package main

import "unicode"

func isPalindrome(s string) bool {
	chars := []rune(s)

	p1 := 0
	p2 := len(chars) - 1

	for p1 < p2 {

		if !unicode.IsNumber(chars[p1]) && !unicode.IsLetter(chars[p1]) {
			p1++
		} else if !unicode.IsNumber(chars[p2]) && !unicode.IsLetter(chars[p2]) {
			p2--
		} else if unicode.ToLower(chars[p1]) == unicode.ToLower(chars[p2]) {
			p1++
			p2--
		} else {
			return false
		}
	}

	return true

}
