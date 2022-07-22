func gameOfLife(board [][]int) {
	directions := [][]int{
		[]int{-1, -1},
		[]int{-1, 0},
		[]int{-1, 1},
		[]int{0, -1},
		[]int{0, 1},
		[]int{1, -1},
		[]int{1, 0},
		[]int{1, 1},
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			count := 0
			for _, dir := range directions {
				count += countLive(board, i+dir[0], j+dir[1])
			}
			cellType := board[i][j] & 1
			if cellType == 0 && count == 3 {
				board[i][j] = 2 // 10
			} else if cellType == 1 && (count == 2 || count == 3) {
				board[i][j] = 3 // 11
			}
		}
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			board[i][j] = board[i][j] >> 1
		}
	}
}

func countLive(board [][]int, i, j int) int {
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[i]) {
		return 0
	}
	return board[i][j] & 1
}
