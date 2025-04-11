package main

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	// for each index
	// check using two pointers in two direction
	ret := ""
	var fn func(i, j int)
	fn = func(i, j int) {
		for i >= 0 && j < len(s) {
			if s[i] != s[j] {
				break
			}
			if len(s[i:j+1]) > len(ret) {
				ret = s[i : j+1]
			}
			i--
			j++
		}
	}
	for i := 0; i < len(s); i++ {
		fn(i, i)
		fn(i, i+1)
	}
	return ret
}
