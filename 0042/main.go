package main

func trap(height []int) int {

	res := 0
	left, right := 0, len(height)-1
	maxLeft, maxRight := 0, 0

	for left < right {
		if height[left] < height[right] {
			maxLeft = Max(maxLeft, height[left])
			res += maxLeft - height[left]
			left++
		} else {
			maxRight = Max(maxRight, height[right])
			res += maxRight - height[right]
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
