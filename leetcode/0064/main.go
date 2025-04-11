package main

import "math"

/*
 * Toolbox: Dynamic Programming
 */

/*
 * Recursive top-down
 */
func minPathSum1(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])

	var fn func(r, c int) int
	fn = func(r, c int) int {
		// base cases
		if r >= rows || c >= cols {
			return math.MaxInt
		}
		if r == rows-1 && c == cols-1 {
			return grid[r][c]
		}
		// recursive
		return grid[r][c] + min(fn(r+1, c), fn(r, c+1))
	}
	return fn(0, 0)
}

/*
/* Recursive top-down with memo
*/
func minPathSum2(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])

	memo := map[[2]int]int{}
	var fn func(r, c int) int
	fn = func(r, c int) int {
		if val, ok := memo[[2]int{r, c}]; ok {
			return val
		}
		// base cases
		if r >= rows || c >= cols {
			return math.MaxInt
		}
		if r == rows-1 && c == cols-1 {
			return grid[r][c]
		}
		// recursive
		ret := grid[r][c] + min(fn(r+1, c), fn(r, c+1))
		memo[[2]int{r, c}] = ret
		return ret
	}
	return fn(0, 0)
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
