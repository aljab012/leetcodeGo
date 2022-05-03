package main

import (
	"fmt"
)

// time: O(n)
// space: O(n)
// Idea: use stack. If you see a closing paren, then that should match the top of the stack.
func isValid(s string) bool {

	chars := []rune(s)

	cMap := make(map[rune]rune)
	cMap[')'] = '('
	cMap['}'] = '{'
	cMap[']'] = '['

	var stack []rune
	for _, char := range chars {
		if char == '(' || char == '{' || char == '[' {
			stack = append(stack, char)
		} else {
			if len(stack) > 0 {
				popChar := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if popChar != cMap[char] {
					return false
				}
			} else {
				return false
			}

		}

	}
	return len(stack) == 0
}
func main() {
	ret := isValid("({}){}")
	fmt.Println(ret)

}
