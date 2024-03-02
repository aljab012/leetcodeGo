package main

import "strings"

/*
 * Recursive top-down
 */
func wordBreak1(s string, wordDict []string) []string {
	ret := []string{}
	var fn func(string, string)
	fn = func(str string, out string) {
		if len(str) == 0 {
			ret = append(ret, out[:len(out)-1])
			return
		}
		for _, w := range wordDict {
			if strings.HasPrefix(str, w) {
				fn(str[len(w):], out+w+" ")
			}
		}
	}
	fn(s, "")
	return ret
}
