package main

import "fmt"

// idea: sliding window technnique
func lengthOfLongestSubstring(s string) int {
	ret := ""

	set := make([]bool, 128)

	left := 0
	for right := 0; right < len(s); right++ {
		for set[s[right]] {
			set[s[left]] = false
			left++
		}
		set[s[right]] = true
		window := s[left : right+1]
		if len(window) > len(ret) {
			ret = window
		}
	}
	return len(ret)
}

func main() {
	// Output: 3
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}
