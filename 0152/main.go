package main

/*
 * Brute force approach with 3 nested loops
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
 */
func maxProduct(nums []int) int {
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
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
