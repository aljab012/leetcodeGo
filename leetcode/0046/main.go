package main

func permute(nums []int) [][]int {
	return helper([]int{}, nums, [][]int{})
}

func helper(cur []int, left []int, res [][]int) [][]int {
	if len(left) == 0 {
		res = append(res, cur)
		return res
	}

	for i, l := range left {
		newCur := append(cur, l)
		newLeft := append(append([]int{}, left[:i]...), left[i+1:]...)
		res = helper(newCur, newLeft, res)
	}
	return res
}
