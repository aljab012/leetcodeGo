package main

import "strconv"

func restoreIpAddresses(s string) []string {
	if len(s) > 12 {
		return []string{}
	}
	return helper(0, 0, "", s, []string{})
}
func helper(i int, dots int, curIP string, s string, ret []string) []string {
	if dots == 4 && i == len(s) {
		ret = append(ret, curIP[:len(curIP)-1])
		return ret
	}
	if dots > 4 {
		return ret
	}
	for j := i; j < Min(i+3, len(s)); j++ {
		n, _ := strconv.Atoi(s[i : j+1])
		if n <= 255 && (i == j || s[i] != '0') {
			ret = helper(i+1, dots+1, curIP+strconv.Itoa(n)+".", s, ret)
		}
	}
	return ret
}
func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
