package main

/*
 * Approach:
 * ToolBox: Binary Search, Matrix
 * 1. Treat the 2D matrix as a 1D array.
 * 2. Use binary search to find the target.
 * Time Complexity: O(log(m * n))
 * Space Complexity: O(1)
 */
func searchMatrix(matrix [][]int, target int) bool {
	rows := len(matrix)
	cols := len(matrix[0])

	l, r := 0, rows*cols-1
	for l <= r {
		mid := (r + l) / 2
		if target == matrix[mid/cols][mid%cols] {
			return true
		} else if target > matrix[mid/cols][mid%cols] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return false
}
