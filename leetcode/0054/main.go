package main

func spiralOrder(matrix [][]int) []int {
	isValid := func(r, c int) bool {
		if r < 0 || c < 0 || r >= len(matrix) || c >= len(matrix[0]) || matrix[r][c] == -101 {
			return false
		}
		return true
	}

	dir := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	d := 0

	ret := []int{}
	r, c := 0, 0
	for len(ret) < len(matrix)*len(matrix[0]) {
		ret = append(ret, matrix[r][c])
		matrix[r][c] = -101

		nr, nc := r+dir[d%4][0], c+dir[d%4][1]
		if !isValid(nr, nc) {
			d++
			nr, nc = r+dir[d%4][0], c+dir[d%4][1]
		}
		r, c = nr, nc
	}
	return ret
}
