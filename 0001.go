package main

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
