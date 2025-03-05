package main

// recursive solution
func canVisitAllRooms1(rooms [][]int) bool {
	visited := map[int]bool{}

	var visit func(int)
	visit = func(i int) {
		if _, ok := visited[i]; ok {
			return
		}
		visited[i] = true
		for _, n := range rooms[i] {
			visit(n)
		}
	}
	visit(0)

	return len(visited) == len(rooms)
}

// iterative solution
func canVisitAllRooms2(rooms [][]int) bool {
	visited := map[int]bool{}
	stack := []int{0}

	for len(stack) > 0 {
		pop := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if _, ok := visited[pop]; ok {
			continue
		}
		visited[pop] = true
		for _, n := range rooms[pop] {
			stack = append(stack, n)
		}
	}
	return len(visited) == len(rooms)
}
