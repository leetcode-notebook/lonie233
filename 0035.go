package main

func searchInsert(nums []int, target int) int {

	if len(nums) == 0 {
		return 0
	}

	return findIndex(nums, target, 0, len(nums)-1)
}

func findIndex(nums [] int, target int, left int, right int) int {
	if left <= right {
		if target < nums[left] {
			return left
		}
		if target > nums[right] {
			return right + 1
		}
	}

	midIndex := (left + right) / 2
	midVal := nums[midIndex]

	if target == midVal {
		return midIndex
	} else if target > midVal {
		return findIndex(nums, target, midIndex+1, right)
	} else {
		return findIndex(nums, target, left, midIndex-1)
	}

}
