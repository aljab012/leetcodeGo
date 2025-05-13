package main

import (
	"container/heap"
	"math/rand"
)

type MaxHeap []int

func (h MaxHeap) Len() int {
	return len(h)
}
func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MaxHeap) Pop() interface{} {
	old := *h
	length := len(old)
	ret := old[length-1]
	*h = old[0 : length-1]
	return ret
}

func findKthLargest(nums []int, k int) int {
	maxHeap := &MaxHeap{}
	heap.Init(maxHeap)
	for _, num := range nums {
		heap.Push(maxHeap, num)
	}
	var pop int
	for i := 0; i < k; i++ {
		pop = heap.Pop(maxHeap).(int)
	}
	return pop
}

func findKthLargest2(nums []int, k int) int {
	k = len(nums) - k

	return quickSelect(nums, k)
}
func quickSelect(nums []int, k int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	pivotIndex := rand.Intn(len(nums))
	pivotNumber := nums[pivotIndex]
	var smaller, larger []int
	for i, n := range nums {
		if i == pivotIndex {
			continue
		}
		if n > pivotNumber {
			larger = append(larger, n)
		} else {
			smaller = append(smaller, n)
		}
	}
	if k < len(smaller) {
		return quickSelect(smaller, k)
	} else if k == len(smaller) {
		return pivotNumber
	} else {
		return quickSelect(larger, k-len(smaller)-1)
	}
}
