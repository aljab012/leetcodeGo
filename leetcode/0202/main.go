package main

func isHappy(n int) bool {
	nMap := map[int]bool{}

	for n != 1 && !nMap[n] {
		nMap[n] = true
		n = sqrSum(n)
	}
	return n == 1
}
func sqrSum(n int) int {
	sum := 0

	for n != 0 {
		digit := n % 10
		n /= 10
		sum += digit * digit
	}
	return sum
}
