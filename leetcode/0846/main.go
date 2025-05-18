package main

import "sort"

func isNStraightHand(hand []int, groupSize int) bool {
	cardsMap := map[int]int{}
	for _, num := range hand {
		cardsMap[num]++
	}
	sortedCards := []int{}
	sortedCards = append(sortedCards, hand...)

	sort.Ints(sortedCards)

	for _, card := range sortedCards {
		if _, ok := cardsMap[card]; !ok {
			continue
		}
		for i := 0; i < groupSize; i++ {
			if _, ok := cardsMap[card+i]; !ok {
				return false
			}
			cardsMap[card+i]--
			if cardsMap[card+i] == 0 {
				delete(cardsMap, card+i)
			}
		}
	}
	return true
}
