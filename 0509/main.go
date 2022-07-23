func fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	n1 := 0
	n2 := 1
	for i := 2; i <= n; i++ {
		temp := n1 + n2
		n1 = n2
		n2 = temp
	}
	return n2
}

// memo
func fibMemo(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	memo := make([]int, n+1)
	memo[0] = 0
	memo[1] = 1
	for i := 2; i < len(memo); i++ {
		memo[i] = memo[i-1] + memo[i-2]
	}
	return memo[n]

}

// Recursive
func fibR(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibR(n-1) + fibR(n-2)
}
