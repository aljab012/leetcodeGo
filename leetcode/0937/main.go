package main

import (
	"sort"
	"strings"
)

// Idea:
// split letter/digit logs
// sort the letter by log:
// if equal sort by the whole string
func reorderLogFiles(logs []string) []string {
	letters := []string{}
	digits := []string{}

	for _, log := range logs {
		i := strings.Index(log, " ")
		if log[i+1] >= 'a' && log[i+1] <= 'z' {
			letters = append(letters, log)
		} else {
			digits = append(digits, log)
		}
	}
	sort.Slice(letters, func(i, j int) bool {
		iIndex := strings.Index(letters[i], " ")
		jIndex := strings.Index(letters[j], " ")

		if letters[i][iIndex:] == letters[j][jIndex:] {
			return letters[i] < letters[j]
		}
		return letters[i][iIndex:] < letters[j][jIndex:]
	})
	return append(letters, digits...)
}
