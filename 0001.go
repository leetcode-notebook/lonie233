package main

import "fmt"

func twoSum(nums []int, target int) []int {

	otherSide := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		otherSide[target-nums[i]] = i
	}

	result := []int{0, 0}
	for index, value := range nums {
		otherIndex, exist := otherSide[value]
		if exist && index != otherIndex {
			result[0] = index
			result[1] = otherIndex
			break
		}
	}
	return result
}

func main() {
	var balance = []int{1, 2, 3}
	fmt.Println(twoSum(balance, 4))
}
