package main

// solution using recursive dfs
func floodFill1(image [][]int, sr int, sc int, color int) [][]int {
	originalColor := image[sr][sc]
	var dfs func(int, int)
	dfs = func(i, j int) {
		if i < 0 || j < 0 || i >= len(image) || j >= len(image[i]) || image[i][j] != originalColor || image[i][j] == color {
			return
		}
		image[i][j] = color
		dfs(i+1, j)
		dfs(i-1, j)
		dfs(i, j+1)
		dfs(i, j-1)
	}
	dfs(sr, sc)
	return image
}

// solution using iterative dfs
func floodFill2(image [][]int, sr int, sc int, color int) [][]int {
	originalColor := image[sr][sc]
	stack := [][]int{{sr, sc}}

	for len(stack) > 0 {
		i, j := stack[len(stack)-1][0], stack[len(stack)-1][1]
		stack = stack[:len(stack)-1]

		if i < 0 || j < 0 || i >= len(image) || j >= len(image[i]) || image[i][j] != originalColor || image[i][j] == color {
			continue
		}
		image[i][j] = color
		stack = append(stack, []int{i + 1, j}, []int{i - 1, j}, []int{i, j + 1}, []int{i, j - 1})
	}
	return image
}
