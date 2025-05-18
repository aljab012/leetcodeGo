package main

func isAlienSorted(words []string, order string) bool {
	charOrder := map[byte]int{}
	for i := 0; i < len(order); i++ {
		charOrder[order[i]] = i
	}
	var isSorted func(w1, w2 string) bool
	isSorted = func(w1, w2 string) bool {
		p1, p2 := 0, 0
		for p1 < len(w1) && p2 < len(w2) {
			if charOrder[w1[p1]] > charOrder[w2[p2]] {
				return false
			} else if charOrder[w1[p1]] < charOrder[w2[p2]] {
				return true
			} else {
				p1++
				p2++
			}
		}
		return len(w1) <= len(w2)
	}

	for i := 1; i < len(words); i++ {
		w1 := words[i-1]
		w2 := words[i]
		if !isSorted(w1, w2) {
			return false
		}
	}
	return true
}
