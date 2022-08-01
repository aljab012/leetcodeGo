package main

func updateMatrix(mat [][]int) [][]int {
	dir := [][]int{
		[]int{0,-1},
		[]int{0,1},
		[]int{-1,0},
		[]int{1,0},
	}
	queue := [][]int{}

	rows, cols := len(mat), len(mat[0])

	for i :=0; i < rows; i++ {
		for j :=0; j < cols; j++{
			if mat[i][j] == 0 {
				queue  = append(queue, []int{i,j})
			}else{
				mat[i][j] = -1
			}
		}
	}
	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]
		for i := range dir{
			row := item[0] + dir[i][0]
			col := item[1] + dir[i][1]
			if row <0 || col < 0 || row >= rows || col >= cols || mat[row][col] != -1{
				continue
			}
			mat[row][col] = mat[item[0]][item[1]]+1
			queue = append(queue,[]int{row,col})
		}
	}
	return mat
}
