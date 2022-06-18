package main

func rotate(matrix [][]int) {

	for i := 0; i < len(matrix); i++ {
		for j := i; j < len(matrix[i]); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]

		}
	}
	for i := range matrix {
		matrix[i] = Reverse(matrix[i], 0, len(matrix[i])-1)
	}
}

func Swap(nums []int, i, j int) []int {
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
	return nums
}

func Reverse(nums []int, i, j int) []int {

	for i < j {
		Swap(nums, i, j)
		i++
		j--
	}
	return nums
}
