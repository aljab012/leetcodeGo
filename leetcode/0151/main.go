package main

import "strings"

func reverseWords(s string) string {
	arr := strings.Fields(s)
	arr = reverse(arr)
	return strings.Join(arr, " ")
}

func reverse(arr []string) []string {
	l, r := 0, len(arr)-1
	for l < r {
		arr[l], arr[r] = arr[r], arr[l]
		l++
		r--
	}
	return arr
}
