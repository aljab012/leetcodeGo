package main

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	const START, END = 0, 1
	result := [][]int{}
	p1, p2 := 0, 0

	for p1 < len(firstList) && p2 < len(secondList) {
		start1, end1 := firstList[p1][START], firstList[p1][END]
		start2, end2 := secondList[p2][START], secondList[p2][END]

		if end1 < start2 {
			p1++ // period1 is before period2
		} else if end2 < start1 {
			p2++ // period2 is before period1
		} else { // they intersect, move the pointer of the period that ends first
			result = append(result, []int{max(start1, start2), min(end1, end2)})
			if end1 < end2 {
				p1++
			} else {
				p2++
			}
		}
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
