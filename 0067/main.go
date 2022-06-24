package main

import (
	"fmt"
	"strconv"
)

func addBinary(a string, b string) string {
	aPtr := len(a) - 1
	bPtr := len(b) - 1

	carry := 0
	ret := ""
	for aPtr >= 0 || bPtr >= 0 {
		aVal := 0
		if aPtr >= 0 {
			aVal, _ = strconv.Atoi(string(a[aPtr]))
			aPtr--
		}
		bVal := 0
		if bPtr >= 0 {
			bVal, _ = strconv.Atoi(string(b[bPtr]))
			bPtr--
		}
		carry += aVal + bVal
		digit := carry % 2
		carry = carry / 2
		ret = strconv.Itoa(digit) + ret
	}
	if carry != 0 {

		ret = strconv.Itoa(carry) + ret
	}
	return ret
}

func main() {
	fmt.Println(addBinary("11", "1"))
}
