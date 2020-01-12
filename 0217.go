package main

func containsDuplicate(nums []int) bool {
	var resMap = make(map[int]int, len(nums))

	for _,number := range nums {
		if resMap[number] == 0 {
			resMap[number] = 1
		}else {
			return true
		}
	}

	return false
}
