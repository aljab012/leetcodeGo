package main

/*
 * Recursive solution
 */
func isInterleave1(s1 string, s2 string, s3 string) bool {
	var fn func(i, j, k int) bool
	fn = func(i, j, k int) bool {
		if k == len(s3) && i == len(s1) && j == len(s2) {
			return true
		}
		if k >= len(s3) {
			return false
		}
		if i < len(s1) && s1[i] == s3[k] && fn(i+1, j, k+1) {
			return true
		}
		if j < len(s2) && s2[j] == s3[k] && fn(i, j+1, k+1) {
			return true
		}
		return false
	}
	return fn(0, 0, 0)
}

/*
 * Recursive solution with memoization
 */
func isInterleave2(s1 string, s2 string, s3 string) bool {
	type key struct{ i, j, k int }
	memo := map[key]bool{}

	var fn func(i, j, k int) bool
	fn = func(i, j, k int) bool {
		if v, ok := memo[key{i, j, k}]; ok {
			return v
		}
		if k == len(s3) && i == len(s1) && j == len(s2) {
			return true
		}
		if k >= len(s3) {
			return false
		}
		if i < len(s1) && s1[i] == s3[k] && fn(i+1, j, k+1) {
			memo[key{i, j, k}] = true
			return true
		}
		if j < len(s2) && s2[j] == s3[k] && fn(i, j+1, k+1) {
			memo[key{i, j, k}] = true
			return true
		}
		memo[key{i, j, k}] = false
		return false
	}
	return fn(0, 0, 0)
}
