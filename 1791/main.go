package main

func findCenter(edges [][]int) int {
	graph := buildGraph(edges)
	for node, neighbors := range graph {
		if len(neighbors) == len(edges) {
			return node
		}
	}
	return -1
}

func buildGraph(edges [][]int) map[int][]int {
	adjList := map[int][]int{}
	for _, edge := range edges {
		src, dest := edge[0], edge[1]
		adjList[src] = append(adjList[src], dest)
		adjList[dest] = append(adjList[dest], src)
	}
	return adjList
}
