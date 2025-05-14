package main

import "container/heap"

type PointWithDist struct {
	distance int
	point    []int
}
type MinHeap []PointWithDist

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].distance < h[j].distance }
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(PointWithDist))
}
func (h *MinHeap) Pop() interface{} {
	old := *h
	pop := old[len(old)-1]
	*h = old[:len(old)-1]
	return pop
}

func kClosest(points [][]int, k int) [][]int {
	minHeap := &MinHeap{}
	heap.Init(minHeap)

	for _, point := range points {
		distance := point[0]*point[0] + point[1]*point[1]
		heap.Push(minHeap, PointWithDist{distance, point})
	}

	result := [][]int{}
	for len(result) < k {
		pop := heap.Pop(minHeap).(PointWithDist)
		result = append(result, pop.point)
	}

	return result
}
