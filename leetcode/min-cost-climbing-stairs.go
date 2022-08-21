/*
- Problem: https://leetcode.com/problems/min-cost-climbing-stairs/
*/

func minCostClimbingStairs(cost []int) int {

	// Defining a cache array to store min cost to climb stairs from each position
	var cache []int = make([]int, len(cost))

	// Min cost to climb stairs from last two steps will always be it's own cost
	cache[len(cost)-1] = cost[len(cost)-1]
	cache[len(cost)-2] = cost[len(cost)-2]

	for i := len(cost) - 3; i >= 0; i-- {
		// Calculating Min cost to climb stairs from current position
		cache[i] = cost[i] + findMin(cache[i+1], cache[i+2])
	}

	// As we can start either from 0 or 1 index, choosing which ever has lesser cost
	if cache[0] < cache[1] {
		return cache[0]
	} else {
		return cache[1]
	}
}

func findMin(val1 int, val2 int) int {
	if val1 < val2 {
		return val1
	}
	return val2
}
