package main

import "strings"

/*
 * Recursive top-down
 */
func wordBreak1(s string, wordDict []string) bool {
	var fn func(string) bool
	fn = func(str string) bool {
		if len(str) == 0 {
			return true
		}
		for _, w := range wordDict {
			if strings.HasPrefix(str, w) {
				if fn(str[len(w):]) {
					return true
				}
			}
		}
		return false
	}
	return fn(s)
}

/*
 * Recursive top-down with memoization
 */
func wordBreak2(s string, wordDict []string) bool {
	memo := map[string]bool{}
	var fn func(string) bool
	fn = func(str string) bool {
		if len(str) == 0 {
			return true
		}
		if val, ok := memo[str]; ok {
			return val
		}
		for _, w := range wordDict {
			if strings.HasPrefix(str, w) {
				if fn(str[len(w):]) {
					memo[str[len(w):]] = true
					return true
				}
			}
		}
		memo[str] = false
		return false
	}
	return fn(s)
}
