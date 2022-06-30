package main

import "fmt"

func uniquePaths(m int, n int) int {
	if m == 0 || n == 0 {
		return 0
	}
	arr := make([]int, n)
	for i := range arr {
		arr[i] = 1
	}
	for i := 0; i < m-1; i++ {
		for j := len(arr) - 2; j >= 0; j-- {
			arr[j] = arr[j] + arr[j+1]
		}
	}
	return arr[0]
}

func main() {
	// expected: 3
	fmt.Println(uniquePaths(3, 2))
}
