package main

import "math"

/*
 * Recursive top-down
 */
func minPathSum1(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])
	var fn func(x, y int) int
	fn = func(x, y int) int {
		if x == rows || y == cols { // base case (out of bound)
			return int(math.MaxInt32)
		}
		if x == rows-1 && y == cols-1 { // base case (hit goal)
			return grid[x][y]
		}
		return grid[x][y] + min(fn(x+1, y), fn(x, y+1))
	}
	return fn(0, 0)
}

/*
/* Recursive top-down with memo
*/
func minPathSum2(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])
	type key struct {
		x, y int
	}
	memo := map[key]int{}
	var fn func(x, y int) int
	fn = func(x, y int) int {
		if x == rows || y == cols { // base case (out of bound)
			return int(math.MaxInt32)
		}
		if x == rows-1 && y == cols-1 { // base case (hit goal)
			return grid[x][y]
		}
		if val, ok := memo[key{x, y}]; ok {
			return val
		}
		ret := grid[x][y] + min(fn(x+1, y), fn(x, y+1))
		memo[key{x, y}] = ret
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
