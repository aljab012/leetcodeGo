package main

func maxArea(height []int) int {
	maxArea := 0

	left, right := 0, len(height)-1
	for left < right {
		area := min(height[left], height[right]) * (right - left)
		maxArea = max(maxArea, area)

		// always move lower column
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
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
