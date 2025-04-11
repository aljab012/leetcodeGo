package main

/*
 * Optimal solution
 * Time complexity: O(n^2)
 * Space complexity: O(1)
 * Intuition: For each character in the string, we can expand outwards from that character to find all palindromic substrings.
 * Two cases:
 * 1. Odd length palindromes: Expand outwards from the current character to the left and right until the characters are not equal.
 * 2. Even length palindromes: Expand outwards from the current character and the next character to the left and right until the characters are not equal.
 */

func countSubstrings(s string) int {
	count := 0
	var fn func(l, r int)
	fn = func(l, r int) {
		for l <= r && l >= 0 && r < len(s) && s[l] == s[r] {
			count++
			l--
			r++
		}
	}
	for i := range s {
		fn(i, i)
		fn(i, i+1)
	}
	return count
}
