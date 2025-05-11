package main

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}

	maxLength := 0

	for num := range numSet {
		// Skip if not the start of a sequence
		if numSet[num-1] {
			continue
		}

		// Count sequence length
		currentLength := 1
		for numSet[num+1] {
			currentLength++
			num++
		}

		maxLength = max(maxLength, currentLength)
	}

	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
