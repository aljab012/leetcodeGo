package main

func maxArea(height []int) int {

	res := 0
	left, right := 0, len(height)-1
	for left < right {
		area := (right - left) * Min(height[left], height[right])
		res = Max(res, area)

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return res
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
