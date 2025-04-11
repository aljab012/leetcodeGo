package main

func numJewelsInStones(jewels string, stones string) int {
	ret := 0
	set := make(map[rune]bool)
	for _, c := range jewels {
		set[c] = true
	}
	for _, c := range stones {
		if _, ok := set[c]; ok {
			ret++
		}
	}
	return ret
}
