package main

/*
 * Idea:
 * 1. Split the problem into subproblems and solve them recursively
 * 2. Use memoization to avoid recomputation
 * 3. The base case when n is 0 or 1
 */

/*
 * Recursive top-down
 */
func fib1(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib1(n-1) + fib1(n-2)
}

/*
 * Recursive top-down with memo
 */
func fib2(n int) int {
	memo := map[int]int{}

	var fn func(int) int
	fn = func(n int) int {
		if val, ok := memo[n]; ok {
			return val
		}
		if n == 0 {
			return 0
		}
		if n == 1 {
			return 1
		}
		memo[n] = fn(n-1) + fn(n-2)
		return memo[n]
	}
	return fn(n)
}

/*
 * Iterative bottom-up with array
 */
func fib3(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

/*
 * Iterative bottom-up with variables
 */

func fib4(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	prev := 0
	cur := 1
	for i := 2; i <= n; i++ {
		cur, prev = cur+prev, cur
	}
	return cur
}
