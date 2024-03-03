package main

/*
  - Intuition:
  - 1. Draw tree diagram to understand the problem
      The tree diagram is a binary tree with 2 branches (move down or move right)
      Move down is subtracting 1 from m
      Move right is subtracting 1 from n
      The leaf node is when m and n are 1
  - 2. The base case is when m and n are 1 (reached the end)
  - 3. The base case is when m or n is 0 (out of bound)
  - 4. Either move down or move right to reach the end of the grid which is m-1 or n-1
*/

/*
 * Top-down recursive
 * Idea: Use recursion to move down or move right
 * 	 The base case is when m and n are 1 (reached the end)
 * 	 The base case is when m or n is 0 (out of bound)
 * 	 The answer is the number of ways to reach the end of the grid which is m-1 or n-1
 * 	 Time complexity: O(2^(m+n))
 * 	 Space complexity: O(m+n)
 */
func uniquePaths1(m int, n int) int {
	var fn func(int, int) int
	fn = func(x int, y int) int {
		if x == 1 && y == 1 { // base case: reached the end
			return 1
		}
		if x == 0 || y == 0 { // base case: out of bound
			return 0
		}
		return fn(x-1, y) + fn(x, y-1) // recursive call: move down or move right
	}
	return fn(m, n)
}

/*
 * Top-down recursive with memo
 * Idea: Use recursion to move down or move right and store the result in a memoization table
 * 	 The base case is when m and n are 1 (reached the end)
 * 	 The base case is when m or n is 0 (out of bound)
 * 	 The answer is the number of ways to reach the end of the grid which is m-1 or n-1
 * 	 Time complexity: O(m*n)
 * 	 Space complexity: O(m*n)
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
		if x == 1 && y == 1 { // base case: reached the end
			return 1
		}
		if x == 0 || y == 0 { // base case: out of bound
			return 0
		}
		ret := fn(x-1, y) + fn(x, y-1) // recursive call: move down or move right
		memo[key{x, y}] = ret
		return ret
	}
	return fn(m, n)
}

/*
 * Top-down iterative with tabulation
 * Idea: Use a 2D array to store the number of ways to reach each cell
 * 	 The number of ways to reach a cell is the sum of the number of ways to reach the cell above and the cell to the left
 * 	 The base case is the first row and the first column
 * 	 The answer is the number of ways to reach the last cell
 * 	 Time complexity: O(m*n)
 * 	 Space complexity: O(m*n)
 */
func uniquePaths3(m int, n int) int {

	matrix := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		matrix[i] = make([]int, n+1)
	}

	matrix[1][1] = 1 // base case
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			cur := matrix[i][j]
			if i+1 <= m {
				matrix[i+1][j] += cur // move down
			}
			if j+1 <= n {
				matrix[i][j+1] += cur // move right
			}
		}
	}
	return matrix[m][n]

}

func uniquePaths4(m int, n int) int {
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
