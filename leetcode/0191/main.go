package main

func hammingWeight(num uint32) int {
	ret := 0
	for num != 0 {
		ret += int(num & 1)
		num = num >> 1
	}
	return ret
}
