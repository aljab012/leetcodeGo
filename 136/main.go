package main

func singleNumber(nums []int) int {
	ret := 0
	for _, n := range nums {
		ret ^= n
	}
	return ret
}
