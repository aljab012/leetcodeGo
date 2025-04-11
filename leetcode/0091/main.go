package main

import (
	"strconv"
)

/*
 * Idea:
 * 1. Split the problem into subproblems and solve them recursively
 * 2. Use memoization to avoid recomputation
 * 3. The base case is when the position is at the end of the string
 * 4. If the current position is 0, then there is no way to decode it
 * 5. If the current position is 1-9, then there is 1 way to decode it
 * 6. If the current position is 10-26, then there is 1 way to decode it
 * 7. The number of ways to decode the string is the sum of the ways to decode
 * the string from the next position and the string from the position after the next position
 */

/*
 * Recursive top-down
 */

func numDecodings1(s string) int {
	var fn func(pos int) int
	fn = func(pos int) int {
		if pos == len(s) {
			return 1
		}
		if s[pos] == '0' {
			return 0
		}
		ways := 0
		if pos+1 <= len(s) {
			if n, _ := strconv.Atoi(s[pos : pos+1]); n >= 1 && n <= 9 {
				ways += fn(pos + 1)
			}
		}
		if pos+2 <= len(s) {
			if n, _ := strconv.Atoi(s[pos : pos+2]); n >= 10 && n <= 26 {
				ways += fn(pos + 2)
			}
		}
		return ways
	}
	return fn(0)
}

/*
 * Recursive top-down with memo
 */
func numDecodings2(s string) int {
	memo := map[int]int{}

	var fn func(int) int
	fn = func(pos int) int {
		if val, ok := memo[pos]; ok {
			return val
		}
		if pos == len(s) {
			return 1
		}
		if s[pos] == '0' {
			return 0
		}
		ways := 0
		if pos+1 <= len(s) {
			d, _ := strconv.Atoi(s[pos : pos+1])
			if d >= 1 && d <= 9 {
				ways += fn(pos + 1)
			}
		}
		if pos+2 <= len(s) {
			dd, _ := strconv.Atoi(s[pos : pos+2])
			if dd >= 10 && dd <= 26 {
				ways += fn(pos + 2)
			}
		}
		memo[pos] = ways
		return ways
	}
	return fn(0)
}

/*
/* Bottom-up iterative with memo
*/
func numDecodings3(s string) int {

	memo := make([]int, len(s)+1)
	memo[0] = 1
	if s[0] != '0' {
		memo[1] = 1
	}
	for i := 2; i <= len(s); i++ {
		d, _ := strconv.Atoi(s[i-1 : i])
		if d >= 1 && d <= 9 {
			memo[i] += memo[i-1]
		}

		dd, _ := strconv.Atoi(s[i-2 : i])
		if dd >= 10 && dd <= 26 {
			memo[i] += memo[i-2]
		}

	}
	return memo[len(s)]
}

/*
 * Bottom-up iterative with variables
 */
func numDecodings4(s string) int {

	one := 1
	two := 0
	if s[0] != '0' {
		two = 1
	}
	for i := 2; i <= len(s); i++ {
		cur := 0
		d, _ := strconv.Atoi(s[i-1 : i])
		if d >= 1 && d <= 9 {
			cur += two
		}

		dd, _ := strconv.Atoi(s[i-2 : i])
		if dd >= 10 && dd <= 26 {
			cur += one
		}
		one = two
		two = cur
	}
	return two
}
