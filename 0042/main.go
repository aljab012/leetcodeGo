package main

/*
 * Brute force solution
 * For each element, find the left and right minimum height and calculate the water
 * Time complexity: O(n^2)
 * Space complexity: O(1)
 */
func trap1(height []int) int {
	ret := 0
	for i := 0; i < len(height); i++ {
		lMax, rMax := 0, 0
		// find the left max
		for j := i; j >= 0; j-- {
			lMax = max(lMax, height[j])
		}
		// find the right max
		for j := i; j < len(height); j++ {
			rMax = max(rMax, height[j])
		}
		// calculate the water bounded by the min of left and right
		ret += min(lMax, rMax) - height[i]
	}
	return ret
}

/*
 * Optimize solution using space to store the left and right max height for each element
 * Time complexity: O(n)
 * Space complexity: O(n)
 */
func trap2(height []int) int {
	ret := 0
	lMax := make([]int, len(height))
	rMax := make([]int, len(height))

	// find the left max for each element
	lMax[0] = height[0]
	for i := 1; i < len(height); i++ {
		lMax[i] = max(lMax[i-1], height[i])
	}
	// find the right max for each element
	rMax[len(height)-1] = height[len(height)-1]
	for i := len(height) - 2; i >= 0; i-- {
		rMax[i] = max(rMax[i+1], height[i])
	}

	// calculate the water for each element which is bounded by the min of left and right
	for i := 0; i < len(height); i++ {
		ret += min(lMax[i], rMax[i]) - height[i]
	}
	return ret
}

/*
 * Optimize solution using two pointers
 * The water is always bounded by the min of left and right
 * The intuition is to use two pointers to track the left and right max height
 * If the left height is less than the right height, the water is bounded by the left height
 * Otherwise, the water is bounded by the right height
 * The water is the min of left and right height minus the current height
 * Time complexity: O(n)
 * Space complexity: O(1)
 */
func trap3(height []int) int {
	ret := 0
	l, r := 0, len(height)-1
	lMax, rMax := 0, 0
	for l < r {
		lMax = max(lMax, height[l])
		rMax = max(rMax, height[r])
		// if the left height is less than the right height, the water is bounded by the left height
		if height[l] < height[r] {
			water := min(lMax, rMax) - height[l]
			if water > 0 {
				ret += water
			}
			l++
		} else {
			// otherwise, the water is bounded by the right height
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
