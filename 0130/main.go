package main

func solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if i == 0 || j == 0 || i == len(board)-1 || j == len(board[i])-1 {
				DFS(board, i, j)
			}
		}
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 'T' {
				board[i][j] = 'O'
			} else {
				board[i][j] = 'X'
			}
		}
	}
}

func DFS(board [][]byte, i, j int) {
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[i]) || board[i][j] != 'O' {
		return
	}
	board[i][j] = 'T'
	DFS(board, i+1, j)
	DFS(board, i-1, j)
	DFS(board, i, j+1)
	DFS(board, i, j-1)
}
