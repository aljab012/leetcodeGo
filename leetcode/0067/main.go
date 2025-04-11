package main

import (
	"fmt"
	"strconv"
)

func addBinary(a string, b string) string {
	ret := ""
	p1 := len(a) - 1
	p2 := len(b) - 1

	carry := 0
	for p1 >= 0 || p2 >= 0 || carry != 0 {
		if p1 >= 0 {
			carry += int(a[p1] - '0')
		}
		if p2 >= 0 {
			carry += int(b[p2] - '0')
		}
		ret = strconv.Itoa(carry%2) + ret
		carry /= 2
		p1--
		p2--
	}
	return ret
}

func main() {
	fmt.Println(addBinary("11", "1"))
}
