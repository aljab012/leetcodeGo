package main

func longestPalindrome(s string) string {
	res := ""
	for i := range s {
		s1 := helper(i, i, s)
		if len(s1) > len(res) {
			res = s1
		}
		s2 := helper(i, i+1, s)
		if len(s2) > len(res) {
			res = s2
		}

	}
	return res
}
func helper(l, r int, s string) string {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return s[l+1 : r]
}

func main() {

}
