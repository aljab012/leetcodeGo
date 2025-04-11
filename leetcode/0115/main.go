package main

/*
 * Recursive solution
 */
func numDistinct1(s string, t string) int {
	var fn func(i int, j int) int
	fn = func(i int, j int) int {
		// check base cases (index out of bound)
		if j == len(t) {
			return 1
		}
		if i == len(s) {
			return 0
		}

		ret := 0
		// 1. skip (i)
		ret += fn(i+1, j)

		// 2. match (i)
		if s[i] != t[j] {
			return ret
		}
		return ret + fn(i+1, j+1)
	}
	return fn(0, 0)
}

/*
 * Recursive solution with memoization
 */
func numDistinct2(s string, t string) int {
	type key struct{ i, j int }
	memo := map[key]int{}
	var fn func(i int, j int) int
	fn = func(i int, j int) int {
		if v, ok := memo[key{i, j}]; ok {
			return v
		}
		// check base cases (index out of bound)
		if j == len(t) {
			return 1
		}
		if i == len(s) {
			return 0
		}

		ret := 0
		// 1. skip (i)
		ret += fn(i+1, j)

		// 2. match (i)
		if s[i] == t[j] {
			ret += fn(i+1, j+1)
		}
		memo[key{i, j}] = ret
		return ret
	}
	return fn(0, 0)
}
