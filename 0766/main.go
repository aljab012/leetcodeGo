package main

// idea: for each cell, validate the upper left cell
func isToeplitzMatrix(matrix [][]int) bool {
	rows, cols := len(matrix), len(matrix[0])

	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			if matrix[i][j] != matrix[i-1][j-1] {
				return false
			}
		}
	}
	return true
}
