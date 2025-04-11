package main

func isSubsequence1(s string, t string) bool {
	for i := range t {
		if len(s) == 0 {
			break
		}
		if t[i] == s[0] {
			s = s[1:]
		}
	}
	return len(s) == 0
}
