package main

/*
 * Top-down approach recursive
 */
func minCostClimbingStairs1(cost []int) int {
	var fn func(int) int
	fn = func(index int) int {
		if index >= len(cost) {
			return 0
		}
		return cost[index] + min(fn(index+1), fn(index+2))
	}
	return min(fn(0), fn(1))
}

/*
 * Top-down approach recursive with memoization
 */
func minCostClimbingStairs2(cost []int) int {
	memo := map[int]int{}
	var fn func(int) int
	fn = func(index int) int {
		if val, ok := memo[index]; ok {
			return val
		}
		if index >= len(cost) {
			return 0
		}

		ret := cost[index] + min(fn(index+1), fn(index+2))
		memo[index] = ret
		return ret
	}
	return min(fn(0), fn(1))
}

/*
 * Bottom-up approach iterative with array
 */
func minCostClimbingStairs3(cost []int) int {
	arr := make([]int, len(cost))
	arr[0] = cost[0]
	arr[1] = cost[1]
	for i := 2; i < len(cost); i++ {
		arr[i] = cost[i] + min(arr[i-1], arr[i-2])
	}
	return min(arr[len(arr)-1], arr[len(arr)-2])
}

/*
 * Bottom-up approach iterative with variables
 */
func minCostClimbingStairs4(cost []int) int {
	first := cost[0]
	second := cost[1]
	for i := 2; i < len(cost); i++ {
		second, first = cost[i]+min(first, second), second
	}
	return min(first, second)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
