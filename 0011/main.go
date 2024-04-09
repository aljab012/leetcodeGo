package main

func maxArea(height []int) int {
	ret := 0

	l, r := 0, len(height)-1
	for l < r {
		area := min(height[l], height[r]) * (r - l)
		ret = max(ret, area)
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return ret
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
