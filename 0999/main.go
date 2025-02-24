package main

func numRookCaptures(board [][]byte) int {
	isValid := func(r, c int) bool {
		if r >= 0 && r < len(board) && c >= 0 && c < len(board[r]) {
			return true
		}
		return false
	}

	directions := [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}

	count := 0
	for r := range board {
		for c := range board[r] {
			if board[r][c] == 'R' {
				for _, dir := range directions {
					new_r := r + dir[0]
					new_c := c + dir[1]
					for isValid(new_r, new_c) {
						if board[new_r][new_c] == 'p' {
							count++
							break
						}
						if board[new_r][new_c] != '.' {
							break
						}
						new_r += dir[0]
						new_c += dir[1]
					}
				}
				break
			}
		}
	}
	return count
}
