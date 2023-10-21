package main

import "fmt"

func main() {
	cost := []int{10, 15, 20, 10, 54, 77, 25, 12, 6}
	result := minCostClimbingStairs(cost)
	fmt.Println(result)
}

func minCostClimbingStairs(cost []int) int {
	n := len(cost)

	minCost := make([]int, n)
	minCost[0] = cost[0]
	minCost[1] = cost[1]

	for i := 2; i < n; i++ {
		minCost[i] = cost[i] + min(minCost[i-1], minCost[i-2])
	}

	return min(minCost[n-1], minCost[n-2])
}

func min(a, b int) int {

	fmt.Println(a, b)
	if a < b {
		return a
	}
	return b
}
