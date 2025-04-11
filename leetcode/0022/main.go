package main

func generateParenthesis(n int) []string {
	return helper("", 0, 0, n, []string{})
}

func helper(s string, open int, close, n int, ret []string) []string {
	if len(s) == n*2 {
		ret = append(ret, s)
	}
	if open < n {
		ret = helper(s+"(", open+1, close, n, ret)
	}
	if close < open {
		ret = helper(s+")", open, close+1, n, ret)
	}
	return ret
}
