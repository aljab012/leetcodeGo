package main

/*
 * Recursive top-down
 */
func uniquePathsWithObstacles1(obstacleGrid [][]int) int {
	rows := len(obstacleGrid)
	cols := len(obstacleGrid[0])
	// top-down recursive
	var fn func(x int, y int) int
	fn = func(x int, y int) int {
		if x == rows || y == cols { // base case (out of bound)
			return 0
		}
		if obstacleGrid[x][y] == 1 { // base case (obstacle)
			return 0
		}
		if x == rows-1 && y == cols-1 { // base case (reach goal)
			return 1
		}

		return fn(x+1, y) + fn(x, y+1)

	}
	return fn(0, 0)
}

/*
 * Recursive top-down with memo
 */

func uniquePathsWithObstacles2(obstacleGrid [][]int) int {
	rows := len(obstacleGrid)
	cols := len(obstacleGrid[0])
	// top-down recursive with memo
	type key struct {
		x int
		y int
	}
	memo := map[key]int{}
	var fn func(x int, y int) int
	fn = func(x int, y int) int {
		if x == rows || y == cols { // base case (out of bound)
			return 0
		}
		if obstacleGrid[x][y] == 1 { // base case (obstacle)
			return 0
		}
		if x == rows-1 && y == cols-1 { // base case (reach goal)
			return 1
		}

		if val, ok := memo[key{x, y}]; ok {
			return val
		}

		ret := fn(x+1, y) + fn(x, y+1)
		memo[key{x, y}] = ret
		return ret

	}
	return fn(0, 0)
}
