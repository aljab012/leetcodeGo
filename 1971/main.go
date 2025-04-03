package main

func validPath(n int, edges [][]int, source int, destination int) bool {
	graph := buildGraph(edges)

	visited := map[int]bool{}
	queue := []int{source}

	for len(queue) > 0 {
		pop := queue[0]
		queue = queue[1:]

		if pop == destination {
			return true
		}

		if _, ok := visited[pop]; ok {
			continue
		}
		visited[pop] = true

		for _, neighbor := range graph[pop] {
			queue = append(queue, neighbor)
		}
	}
	return false

}
func buildGraph(edges [][]int) map[int][]int {
	ret := map[int][]int{}
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		ret[a] = append(ret[a], b)
		ret[b] = append(ret[b], a)
	}
	return ret
}
