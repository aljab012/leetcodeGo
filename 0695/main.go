func maxAreaOfIsland(grid [][]int) int {
	maxArea := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				maxArea = max(maxArea, DFS(grid, i, j))
			}
		}
	}
	return maxArea
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func DFS(grid [][]int, i, j int) int {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[i]) || grid[i][j] != 1 {
		return 0
	}
	grid[i][j] = 0
	return 1 + DFS(grid, i, j-1) + DFS(grid, i, j+1) + DFS(grid, i-1, j) + DFS(grid, i+1, j)

}
