package main

import "strings"

func reverseWords(s string) string {
	words := strings.Split(s, " ")
	for i := range words {
		words[i] = reverseStr(words[i])
	}
	return strings.Join(words, " ")
}

func reverseStr(s string) string {

	arr := []rune(s)
	l := 0
	r := len(s) - 1

	for l < r {
		arr[l], arr[r] = arr[r], arr[l]
		l++
		r--
	}
	return string(arr)
}
