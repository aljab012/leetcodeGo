package main

func hammingDistance(x int, y int) int {
	ret := 0
	diff := x ^ y

	for diff != 0 {
		if diff&1 != 0 {
			ret++
		}
		diff = diff >> 1
	}
	return ret
}
