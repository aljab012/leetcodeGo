package main

// idea: use set to look up values in O(1)
func longestConsecutive(nums []int) int {
	set := make(map[int]bool)
	max := 0
	for _, n := range nums {
		set[n] = true
	}
	for _, n := range nums {
		if !set[n-1] {
			seqLen := 1
			next := n + 1
			for set[next] {
				seqLen++
				next++
			}
			max = Max(max, seqLen)
		}
	}
	return max
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
