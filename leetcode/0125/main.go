package main

import "unicode"

/*
 * Two pointers approach
 * Time complexity: O(n)
 * Space complexity: O(1)
 */
func isPalindrome(s string) bool {
	chars := []rune(s)

	l := 0
	r := len(chars) - 1

	for l < r {

		if !unicode.IsNumber(chars[l]) && !unicode.IsLetter(chars[l]) {
			l++
		} else if !unicode.IsNumber(chars[r]) && !unicode.IsLetter(chars[r]) {
			r--
		} else if unicode.ToLower(chars[l]) == unicode.ToLower(chars[r]) {
			l++
			r--
		} else {
			return false
		}
	}

	return true

}
