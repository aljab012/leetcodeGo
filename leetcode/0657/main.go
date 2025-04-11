package main

func judgeCircle(moves string) bool {
	count := make([]int, 2)
	for i := range moves {
		if moves[i] == 'U' {
			count[0]++
		} else if moves[i] == 'D' {
			count[0]--
		} else if moves[i] == 'R' {
			count[1]++
		} else if moves[i] == 'L' {
			count[1]--
		}
	}
	return count[0] == 0 && count[1] == 0
}
