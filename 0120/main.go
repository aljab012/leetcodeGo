package main

func minimumTotal(triangle [][]int) int {
	rows := len(triangle)
	memo := map[[2]int]int{}

	var dfs func(r, c int) int
	dfs = func(r, c int) int {
		if val, ok := memo[[2]int{r, c}]; ok {
			return val
		}
		if r == rows-1 {
			return triangle[r][c]
		}
		minPath := triangle[r][c] + min(dfs(r+1, c), dfs(r+1, c+1))
		memo[[2]int{r, c}] = minPath
		return minPath
	}

	return dfs(0, 0)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
