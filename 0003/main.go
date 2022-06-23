package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	set := [128]bool{}
	length, max := 0, 0

	for start, end := 0, 0; end < len(s); end++ {
		char := s[end]

		for ; set[char]; start++ {
			set[s[start]] = false
			length--
		}

		set[char] = true
		length++
		if length > max {
			max = length
		}
	}
	return max
}

func main() {
	// Output: 3
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}
