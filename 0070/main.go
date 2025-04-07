package main

/*
 * Top-down approach recursive
 * Idea: Draw the recursion tree then solve the problem
 */
func climbStairs1(n int) int {
	var fn func(int) int
	fn = func(num int) int {
		if num == 0 {
			return 1
		}
		if num < 0 {
			return 0
		}
		return fn(num-1) + fn(num-2)
	}
	return fn(n)
}

/*
 * Top-down approach recursive with memoization
 * Idea: Use a map to store results of subproblems
 */
func climbStairs2(n int) int {
	memo := map[int]int{}
	var fn func(int) int
	fn = func(num int) int {
		if val, ok := memo[num]; ok {
			return val
		}
		if num == 0 {
			return 1
		}
		if num < 0 {
			return 0
		}
		ret := fn(num-1) + fn(num-2)
		memo[num] = ret
		return ret
	}
	return fn(n)
}

/*
 * Bottom-up approach iterative with array
 * Idea: Use an array to mock fuction calls. Start from the base cases till you reach the top initial problem.
 */
func climbStairs3(n int) int {
	arr := make([]int, n+1)
	arr[0] = 1
	arr[1] = 1
	for i := 2; i <= n; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}
	return arr[n]
}

/*
 * Bottom-up approach iterative with variables
 * Idea: From the previous solution, we can see that we only need the last two values to calculate the next value.
 * So we can use two variables to store the last two values instead of using an array.
 */
func climbStairs4(n int) int {
	first := 1
	second := 1
	for i := 2; i <= n; i++ {
		second, first = first+second, second
	}
	return second
}
