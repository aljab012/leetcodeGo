package main

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	oldColor := image[sr][sc]
	if oldColor == newColor {
		return image
	}
	DFS(image, sr, sc, oldColor, newColor)
	return image
}

func DFS(image [][]int, i, j, oldColor, newColor int) {
	if i < 0 || j < 0 || i >= len(image) || j >= len(image[0]) || image[i][j] != oldColor {
		return
	}
	image[i][j] = newColor
	DFS(image, i+1, j, oldColor, newColor)
	DFS(image, i-1, j, oldColor, newColor)
	DFS(image, i, j+1, oldColor, newColor)
	DFS(image, i, j-1, oldColor, newColor)
}
