package main

import "strconv"

func thousandSeparator(n int) string {
	str := strconv.Itoa(n)
	for i := len(str) - 3; i > 0; i -= 3 {
		str = str[:i] + "." + str[i:]
	}
	return str
}
