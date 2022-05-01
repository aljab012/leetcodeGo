package main

import "fmt"

// time: O(n)
// space: O (n)
func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i, n := range nums {
		comp := target - n
		number, found := numMap[comp]
		if found {
			return []int{number, i}
		}
		numMap[n] = i
	}
	return []int{-1, -1}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	ret := twoSum(nums, target)

	fmt.Println(ret)

}

// 1. brute force solution: double for loops. O(n2)
//
// 2. two pass solution: interate through the array and save num, index into a map.
// then iterate second time while checking for complement. O(n)
//
// 3. one pass solution: implemented above.

// Idea: interate through the array, see if you have encountered a complement in the previous nums.
