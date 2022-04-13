package main

import (
	"fmt"
	"strconv"
)

// time: O(n)
// space: O(n)
func isPalindrome(x int) bool {

	str := strconv.Itoa(x)

	chars := []rune(str)

	p1 := 0
	p2 := len(chars) - 1

	for p1 < p2 {
		if chars[p1] != chars[p2] {
			return false
		}
		p1 += 1
		p2 -= 1
	}
	return true

}

func main() {
	ret := isPalindrome(121)
	fmt.Println(ret)

}
