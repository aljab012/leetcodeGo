package main

// idea: use a map to build buttom up solution
func countPrimes(n int) int {
	ret := 0
	if n <= 2 {
		return 0
	}
	prime := make([]bool, n)
	for i := range prime {
		prime[i] = true
	}
	prime[2] = true
	for i := 2; i < n; i++ {
		if prime[i] {
			ret++
			for j := i; j < n; j += i {
				prime[j] = false
			}
		}
	}
	return ret
}
