package main

/*
 * Brute force approach with 3 nested loops
 * The idea is to generate all subarrays and calculate the product of each subarray
 * Time complexity: O(n^3)
 * Space complexity: O(1)
 */
func maxProduct1(nums []int) int {
	maxProd := nums[0]
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			curProd := 1
			for k := i; k <= j; k++ {
				curProd *= nums[k]
			}
			maxProd = max(maxProd, curProd)
		}
	}
	return maxProd
}

/*
 * Brute force approach with 2 nested loops
 * The idea is to generate all subarrays and calculate the product of each subarray
 * Time complexity: O(n^2)
 * Space complexity: O(1)
 */
func maxProduct2(nums []int) int {
	maxProd := nums[0]
	for i := 0; i < len(nums); i++ {
		curProd := 1
		for j := i; j < len(nums); j++ {
			curProd *= nums[j]
			maxProd = max(maxProd, curProd)
		}
	}
	return maxProd
}

/*
 * Dynamic programming approach using 2 variables
 * The idea is to keep track of the maximum and minimum product of the subarrays ending at the current element
 * This is derived from the drawing of the brute force approach with 2 nested loops from the back
 * and the observation that the maximum product of the subarray ending at the current element can be obtained by multiplying the current element with the maximum product of the subarray ending at the previous element
 * Time complexity: O(n)
 * Space complexity: O(1)
 */
func maxProduct(nums []int) int {
	globalMax := nums[0]
	localMax := 1
	localMin := 1
	for _, n := range nums {
		temp := localMax
		localMax = max(n, max(localMax*n, localMin*n))
		localMin = min(n, min(localMin*n, temp*n))
		globalMax = max(globalMax, localMax)
	}
	return globalMax
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
