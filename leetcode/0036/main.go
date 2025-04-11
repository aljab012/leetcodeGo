package main

func isValidSudoku(board [][]byte) bool {
	// validate rows
	for r := 0; r < len(board); r++ {
		set := map[byte]bool{}
		for c := 0; c < len(board[0]); c++ {
			if _, exist := set[board[r][c]]; exist {
				return false
			}
			if board[r][c] == '.' {
				continue
			}
			set[board[r][c]] = true
		}
	}
	// validate columns
	for c := 0; c < len(board[0]); c++ {
		set := map[byte]bool{}
		for r := 0; r < len(board); r++ {
			if _, exist := set[board[r][c]]; exist {
				return false
			}
			if board[r][c] == '.' {
				continue
			}
			set[board[r][c]] = true
		}
	}
	// validate sub-grids
	for r := 0; r < len(board); r += 3 {
		for c := 0; c < len(board[0]); c += 3 {
			set := map[byte]bool{}
			for subR := r; subR < r+3; subR++ {
				for subC := c; subC < c+3; subC++ {
					if _, exist := set[board[subR][subC]]; exist {
						return false
					}
					if board[subR][subC] == '.' {
						continue
					}
					set[board[subR][subC]] = true
				}
			}
		}
	}
	return true
}
