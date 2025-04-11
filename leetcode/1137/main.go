package main

// tribonacci using naive recursion.
func tribonacci1(n int) int {
	var fn func(int) int
	fn = func(n int) int {
		if n <= 0 {
			return 0
		}
		if n == 1 || n == 2 {
			return 1
		}
		return fn(n-1) + fn(n-2) + fn(n-3)
	}
	return fn(n)
}

// tribonacci using recursion with memoization.
func tribonacci2(n int) int {
	memo := make(map[int]int)
	var fn func(int) int
	fn = func(n int) int {
		if v, ok := memo[n]; ok {
			return v
		}
		if n <= 0 {
			return 0
		}
		if n == 1 || n == 2 {
			return 1
		}
		ret := fn(n-1) + fn(n-2) + fn(n-3)
		memo[n] = ret
		return ret
	}
	return fn(n)
}

// tribonacci using dynamic programming with an array.
func tribonacci3(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}

	arr := make([]int, n+1)
	arr[0], arr[1], arr[2] = 0, 1, 1
	for i := 3; i <= n; i++ {
		arr[i] = arr[i-1] + arr[i-2] + arr[i-3]
	}
	return arr[n]
}

// tribonacci using an iterative approach with constant space.
func tribonacci4(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}

	a, b, c := 0, 1, 1
	for i := 3; i <= n; i++ {
		a, b, c = b, c, a+b+c
	}
	return c
}
