package main

func sortArray(nums []int) []int {
	return bubble(nums)
}

func bubble(nums []int) []int {
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				tmp := nums[i]
				nums[i] = nums[j]
				nums[j] = tmp
			}
		}
	}

	return nums
}
