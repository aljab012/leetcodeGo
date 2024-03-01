package main

/*
  - Intuition:
  - 1. Draw tree diagram to understand the problem
  - 2. The base case is when m and n are 1 (reached the end)
  - 3. The base case is when m or n is 0 (out of bound)
  - 4. Either move down or move right to reach the end of the grid which is m-1 or n-1
*/

/*
 * Top-down recursive
 */
func uniquePaths1(m int, n int) int {
	var fn func(int, int) int
	fn = func(x int, y int) int {
		if x == 1 && y == 1 { // base case reached the end
			return 1
		}
		if x == 0 || y == 0 { // base case out of bound
			return 0
		}
		return fn(x-1, y) + fn(x, y-1) // move down or move right
	}
	return fn(m, n)
}

/*
 * Top-down recursive with memo
 */
func uniquePaths2(m int, n int) int {
	type key struct {
		x int
		y int
	}
	memo := map[key]int{}
	var fn func(int, int) int
	fn = func(x int, y int) int {
		if val, ok := memo[key{x, y}]; ok {
			return val
		}
		if x == 1 && y == 1 {
			return 1
		}
		if x == 0 || y == 0 {
			return 0
		}
		ret := fn(x-1, y) + fn(x, y-1)
		memo[key{x, y}] = ret
		return ret
	}
	return fn(m, n)
}
func uniquePaths3(m int, n int) int {
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
