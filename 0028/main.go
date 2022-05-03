package main

import (
	"fmt"
)

// idea: loop from the begining, if you see needle ahead, return it.
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}

func main() {
	haystack := "hello"
	needle := "ll"

	fmt.Println(strStr(haystack, needle))
}
