package main

// idea: check the slope between points
func checkStraightLine(coord [][]int) bool {
	if len(coord) <= 1 {
		return true
	}
	x, y := 0, 1
	dx := coord[1][x] - coord[0][x]
	dy := coord[1][y] - coord[0][y]

	for i := 1; i < len(coord); i++ {
		ddx := coord[i][x] - coord[i-1][x]
		ddy := coord[i][y] - coord[i-1][y]
		if dx*ddy != dy*ddx {
			return false
		}
	}
	return true
}
