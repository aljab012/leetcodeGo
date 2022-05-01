package main

import "fmt"

// time: O(n*m)
// space: O (n)
func longestCommonPrefix(strs []string) string {
	ret := ""
	if len(strs) == 0 {
		return ret
	}
	for firstStrIndex, _ := range strs[0] {
		for otherIndex, _ := range strs {
			if firstStrIndex == len(strs[otherIndex]) || strs[0][firstStrIndex] != strs[otherIndex][firstStrIndex] {
				return ret
			}

		}
		ret = ret + string(strs[0][firstStrIndex])
	}
	return ret
}

func main() {
	strs := []string{"flowers", "flow", "flight"}

	ret := longestCommonPrefix(strs)

	fmt.Println(ret)

}

// Idea: Take all chars from the first string. Compare that with other strings, if they are not equal or we are out of bound. return what we have.
