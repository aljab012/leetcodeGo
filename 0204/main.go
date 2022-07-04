package main

func countPrimes(n int) int {
	ret := 0
	if n <= 2 {
		return ret
	}
	notPrime := make([]bool, n)
	for i := 2; i < n; i++ {
		if !notPrime[i] {
			ret++
			for j := i; j < n; j += i {
				notPrime[j] = true
			}
		}
	}
	return ret
}
