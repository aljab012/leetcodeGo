package main

/*
 * Recursive solution
 */
func longestIncreasingPath1(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	var fn func(i, j int) int
	fn = func(i, j int) int {
		ret := 1
		dir := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
		for _, d := range dir {
			x := i + d[0]
			y := j + d[1]
			if x >= 0 && y >= 0 && x < len(matrix) && y < len(matrix[x]) && matrix[x][y] > matrix[i][j] {
				ret = max(ret, 1+fn(x, y))
			}
		}
		return ret
	}
	ret := 1
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			ret = max(ret, fn(i, j))
		}
	}
	return ret
}

/*
 * Recursive solution with memoization
 */
func longestIncreasingPath2(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	type key struct{ i, j int }
	memo := map[key]int{}
	var fn func(i, j int) int
	fn = func(i, j int) int {
		if v, ok := memo[key{i, j}]; ok {
			return v
		}
		ret := 1
		dir := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
		for _, d := range dir {
			x := i + d[0]
			y := j + d[1]
			if x >= 0 && y >= 0 && x < len(matrix) && y < len(matrix[x]) && matrix[x][y] > matrix[i][j] {
				ret = max(ret, 1+fn(x, y))
			}
		}
		memo[key{i, j}] = ret
		return ret
	}
	ret := 1
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			ret = max(ret, fn(i, j))
		}
	}
	return ret
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
