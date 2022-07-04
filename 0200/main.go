package main

func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	ret := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				grid = DFS(i, j, grid)
				ret++
			}
		}
	}
	return ret
}
func DFS(i, j int, grid [][]byte) [][]byte {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[i]) || grid[i][j] != '1' {
		return grid
	}
	grid[i][j] = '0'
	grid = DFS(i+1, j, grid)
	grid = DFS(i-1, j, grid)
	grid = DFS(i, j+1, grid)
	grid = DFS(i, j-1, grid)
	return grid
}
