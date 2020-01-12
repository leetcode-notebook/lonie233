package main

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	slow := 0
	for _,value := range nums {
		if value != val {
			nums[slow] = value
			slow++
		}
	}
	return slow
}
