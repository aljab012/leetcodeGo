package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	ROWS, COLS := len(matrix), len(matrix[0])
	top, bot := ROWS-1, 0
	for bot <= top {
		row := bot + (top-bot)/2
		if target > matrix[row][COLS-1] {
			bot = row + 1
		} else if target < matrix[row][0] {
			top = row - 1
		} else {
			break
		}
	}
	if bot > top {
		return false
	}
	targetRow := bot + (top-bot)/2
	left, right := 0, COLS-1
	for left <= right {
		mid := left + (right-left)/2
		if target > matrix[targetRow][mid] {
			left = mid + 1
		} else if target < matrix[targetRow][mid] {
			right = mid - 1
		} else {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println(searchMatrix([][]int{{1}, {3}}, 3))
}
