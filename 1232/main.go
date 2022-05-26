package main

func checkStraightLine(coordinates [][]int) bool {
	if len(coordinates) < 2 {
		return true
	}
	x1 := coordinates[0][0]
	y1 := coordinates[0][1]
	x2 := coordinates[1][0]
	y2 := coordinates[1][1]

	dx := x2 - x1
	dy := y2 - y1

	for i := 2; i < len(coordinates); i++ {

		if dx*(coordinates[i][1]-y2) != dy*(coordinates[i][0]-x2) {
			return false
		}
	}
	return true
}
