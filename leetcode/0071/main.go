package main

import (
	"fmt"
	"strings"
)

func simplifyPath(path string) string {

	stack := []string{}
	for _, chunk := range strings.Split(path, "/") {
		switch chunk {
		case ".", "":
			continue
		case "..":
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		default:
			stack = append(stack, chunk)
		}
	}
	return "/" + strings.Join(stack, "/")
}

func main() {

	// expected: /home/foo
	fmt.Println(simplifyPath("/home//foo"))

}
