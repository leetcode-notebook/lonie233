package main

func minCostClimbingStairs(cost []int) int {

	if cost == nil {
		return 0
	}

	if len(cost) < 2 {
		return cost[0]
	}

	var res = make([]int, len(cost))
	res[0] = cost[0]
	res[1] = cost[1]
	for index, curCost := range cost {
		if index < 2 {
			continue
		}

		res[index] = getTwoMinNumber(res[index-1]+curCost, res[index-2]+curCost)
	}
	return getTwoMinNumber(res[len(cost)-2], res[len(cost)-1])
}
