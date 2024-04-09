package main

func trap(height []int) int {
	// water always bound by the min of left and right
	ret := 0
	l, r := 0, len(height)-1
	lMax, rMax := 0, 0
	for l < r {
		lMax = max(lMax, height[l])
		rMax = max(rMax, height[r])
		if height[l] < height[r] {
			water := min(lMax, rMax) - height[l]
			if water > 0 {
				ret += water
			}
			l++
		} else {
			water := min(lMax, rMax) - height[r]
			if water > 0 {
				ret += water
			}
			r--
		}
	}
	return ret
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
