package main

/*
 * Recusive solution
 */
func longestCommonSubsequence1(text1 string, text2 string) int {
	var fn func(s1 string, s2 string) int
	fn = func(s1 string, s2 string) int {
		// base case
		if len(s1) == 0 || len(s2) == 0 {
			return 0
		}
		if s1[len(s1)-1] == s2[len(s2)-1] {
			return 1 + fn(s1[:len(s1)-1], s2[:len(s2)-1])
		}
		// recursive call
		return max(fn(s1[:len(s1)-1], s2), fn(s1, s2[:len(s2)-1]))
	}
	return fn(text1, text2)
}

/*
/* Recusive solution with memoization
*/
func longestCommonSubsequence2(text1 string, text2 string) int {
	type key struct {
		s1, s2 string
	}
	memo := map[key]int{}
	var fn func(s1 string, s2 string) int
	fn = func(s1 string, s2 string) int {
		if val, ok := memo[key{s1, s2}]; ok {
			return val
		}
		// base case
		if len(s1) == 0 || len(s2) == 0 {
			return 0
		}
		if s1[len(s1)-1] == s2[len(s2)-1] {
			return 1 + fn(s1[:len(s1)-1], s2[:len(s2)-1])
		}
		// recursive call
		ret := max(fn(s1[:len(s1)-1], s2), fn(s1, s2[:len(s2)-1]))
		memo[key{s1, s2}] = ret
		return ret
	}
	return fn(text1, text2)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
