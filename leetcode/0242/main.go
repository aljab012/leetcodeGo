package main

import (
	"sort"
	"strings"
)

/*
 * solution using a map to store the number and check if the number is already in the map
 * Space complexity: O(n)
 * Time complexity: O(n)
 */
func isAnagram1(s string, t string) bool {

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

/*
 * solution using sorting
 * Space complexity: O(n)
 * Time complexity: O(nlogn)
 */
func isAnagram2(s string, t string) bool {
	s = sortString(s)
	t = sortString(t)
	return s == t
}
func sortString(s string) string {
	split := strings.Split(s, "")
	sort.Strings(split)
	return strings.Join(split, "")
}
