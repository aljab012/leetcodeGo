package main

func isAnagram(s string, t string) bool {
	arr := [128]int{}

	for _, c := range s {
		arr[c]++
	}
	for _, c := range t {
		arr[c]--
	}
	for _, n := range arr {
		if n != 0 {
			return false
		}
	}
	return true
}
