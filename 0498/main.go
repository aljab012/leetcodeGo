package main

func findDiagonalOrder(mat [][]int) []int {
	isValid := func(r, c int) bool {
		if r >= 0 && c >= 0 && r < len(mat) && c < len(mat[r]) {
			return true
		}
		return false
	}
	dir := [][]int{{-1, 1}, {1, -1}}

	d := 0

	r, c := 0, 0
	ret := []int{}
	rows := len(mat)
	cols := len(mat[0])
	for len(ret) < cols*rows {
		ret = append(ret, mat[r][c])

		nr := r + dir[d%2][0]
		nc := c + dir[d%2][1]

		if isValid(nr, nc) {
			r, c = nr, nc
		} else {
			if d%2 == 0 {
				if c < cols-1 {
					c += 1 // top adjustment
				} else {
					r += 1 // right adjustment
				}
			} else {
				if r < rows-1 {
					r += 1 // left adjustment
				} else {
					c += 1 // buttom adjustment
				}
			}
			d++
		}

	}
	return ret
}
