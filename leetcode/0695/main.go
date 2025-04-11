package main

func maxAreaOfIsland(grid [][]int) int {
	var dfs func(r, c int) int
	dfs = func(r, c int) int {
		if r < 0 || c < 0 || r >= len(grid) || c >= len(grid[r]) || grid[r][c] != 1 {
			return 0
		}
		grid[r][c] = 0
		return 1 + dfs(r+1, c) + dfs(r-1, c) + dfs(r, c+1) + dfs(r, c-1)
	}
	maxArea := 0
	for r := range grid {
		for c := range grid[0] {
			if grid[r][c] == 1 {
				maxArea = max(maxArea, dfs(r, c))
			}
		}
	}
	return maxArea
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
