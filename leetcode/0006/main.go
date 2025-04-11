package main

/*
 * Idea: create a list of strings with the number of rows
 * iterate through the string and add the character to the corresponding row
 * Space complexity: O(n)
 * Time complexity: O(n)
 */

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	rows := make([]string, numRows)

	direction := 1
	row_index := 0

	for i := 0; i < len(s); i++ {
		rows[row_index] = rows[row_index] + string(s[i])
		if row_index == 0 { // if we are at the top row, we need to go down
			direction = 1
		} else if row_index == len(rows)-1 { // if we are at the bottom row, we need to go up
			direction = -1
		}
		row_index += direction
	}
	ret := ""
	for i := 0; i < len(rows); i++ {
		ret += rows[i]
	}
	return ret
}
