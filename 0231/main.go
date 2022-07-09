package main

// 1000000
// &
// 0111111
// =
// 0000000
func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	return n&(n-1) == 0
}
