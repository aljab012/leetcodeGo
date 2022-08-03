package main

func canVisitAllRooms(rooms [][]int) bool {
	visited := make(map[int]bool)
	stack := []int{0}

	for len(stack) > 0 {
		pop := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		visited[pop] = true

		for _, v := range rooms[pop] {
			if _, ok := visited[v]; !ok {
				stack = append(stack, v)
			}
		}
	}
	return len(visited) == len(rooms)
}
