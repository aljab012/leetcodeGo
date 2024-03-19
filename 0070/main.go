package main

/*
 * Top-down approach recursive
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
/* Bottom-up approach iterative with variables
*/
func climbStairs4(n int) int {
	first := 1
	second := 1
	for i := 2; i <= n; i++ {
		second, first = first+second, second
	}
	return second
}
