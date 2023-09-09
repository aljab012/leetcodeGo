package main

// isAnagram checks if two strings are anagrams
// Time Complexity: O(n)
// Space Complexity: O(1)
func isAnagram(s string, t string) bool {

	letters := make([]int, 26)

	for _, c := range s {
		letters[c-'a']++
	}
	for _, c := range t {
		letters[c-'a']--
	}

	for _, l := range letters {
		if l != 0 {
			return false
		}
	}
	return true

}
