package main

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	ret := [][]int{}

	i, j := 0, 0

	for i < len(firstList) && j < len(secondList) {
		start := Max(firstList[i][0], secondList[j][0])
		end := Min(firstList[i][1], secondList[j][1])

		if start <= end {
			ret = append(ret, []int{start, end})
		}
		if firstList[i][1] < secondList[j][1] {
			i++
		} else {
			j++
		}
	}
	return ret
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
