package main

func checkStraightLine(coord [][]int) bool {
	x, y := 0, 1

	if len(coord) <= 1 {
		return true
	}

	dx := coord[0][x] - coord[1][x]
	dy := coord[0][y] - coord[1][y]

	for i := 1; i < len(coord); i++ {
		tdx := coord[i-1][x] - coord[i][x]
		tdy := coord[i-1][y] - coord[i][y]
		if dy*tdx != dx*tdy {
			return false
		}

	}
	return true
}
