package main

import "math"

func minFallingPathSum(matrix [][]int) int {
	rows, cols := len(matrix), len(matrix[0])
	type key struct{ r, c int }
	memo := map[key]int{}

	result := math.MaxInt

	var dfs func(r, c int) int
	dfs = func(r, c int) int {
		if r < 0 || c < 0 || r >= rows || c >= cols {
			return math.MaxInt - 101
		}
		if r == rows-1 {
			return matrix[r][c]
		}
		if val, ok := memo[key{r, c}]; ok {
			return val
		}
		ret := min(matrix[r][c]+dfs(r+1, c-1), min(matrix[r][c]+dfs(r+1, c), matrix[r][c]+dfs(r+1, c+1)))
		memo[key{r, c}] = ret
		return ret
	}
	for i := 0; i < cols; i++ {
		result = min(result, dfs(0, i))
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
