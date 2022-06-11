package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	ret := 0
	for x != 0 {
		//math.MaxInt32 = 2147483647
		//math.MinInt32 = -2147483648
		digit := x % 10
		x /= 10
		if ret > (math.MaxInt32/10) || ((ret == math.MaxInt32/10) && digit >= math.MaxInt32%10) {
			return 0
		}
		if ret < (math.MinInt32/10) || ((ret == math.MinInt32/10) && digit >= math.MaxInt32%10) {
			return 0
		}
		ret = ret*10 + digit
	}
	return ret
}

func main() {
	fmt.Println(reverse(-123))
}
