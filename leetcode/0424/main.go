package main

func characterReplacement(s string, k int) int {
	count := make([]int, 128)
	ret, left, maxf := 0, 0, 0

	for right := 0; right < len(s); right++ {
		count[s[right]]++
		maxf = Max(maxf, count[s[right]])

		for right-left+1-maxf > k {
			count[s[left]]--
			left++
		}
		ret = Max(ret, right-left+1)

	}
	return ret
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
