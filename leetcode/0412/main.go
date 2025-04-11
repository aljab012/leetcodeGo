package main

import "strconv"

func fizzBuzz(n int) []string {
	ret := make([]string, n)
	for i := range ret {
		n := i + 1
		if n%3 == 0 && n%5 == 0 {
			ret[i] = "FizzBuzz"
		} else if n%3 == 0 {
			ret[i] = "Fizz"
		} else if n%5 == 0 {
			ret[i] = "Buzz"
		} else {
			ret[i] = strconv.Itoa(n)
		}
	}
	return ret
}
