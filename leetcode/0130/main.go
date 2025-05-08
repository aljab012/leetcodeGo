package main

func solve(board [][]byte) {
	rows := len(board)
	cols := len(board[0])

	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || c < 0 || r >= rows || c >= cols || board[r][c] != 'O' {
			return
		}
		board[r][c] = 'B'
		dfs(r+1, c)
		dfs(r, c+1)
		dfs(r-1, c)
		dfs(r, c-1)
	}
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if r == 0 || c == 0 || r == rows-1 || c == cols-1 && board[r][c] == 'O' {
				dfs(r, c)
			}
		}
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if board[r][c] == 'B' {
				board[r][c] = 'O'
			} else if board[r][c] == 'O' {
				board[r][c] = 'X'
			}
		}
	}
}
