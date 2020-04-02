package main

func sortArray(nums []int) []int {
	//return bubble(nums)
	//return mergeSort(nums)
	return quickSort(nums)
}

func quickSort(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}

	quickSortReal(nums, 0, len(nums)-1)

	return nums
}

func quickSortReal(nums []int, left int, right int) {

	if left < right {
		index := getIndex(nums, left, right)
		quickSortReal(nums, left, index-1)
		quickSortReal(nums, index+1, right)
	}

}

func getIndex(nums []int, left int, right int) int {

	provt := nums[left]

	for left < right {

		for left < right && nums[right] > provt {
			right--
		}
		if left < right {
			nums[left] = nums[right]
			left++
		}

		for left < right && nums[left] < provt {
			left++
		}

		if left < right {
			nums[right] = nums[left]
			right--
		}

	}

	nums[left] = provt

	return left

}

func swap(nums []int, left int, right int) {
	tmp := nums[left]
	nums[left] = nums[right]
	nums[right] = tmp
}

func mergeSort(nums []int) []int {

	if len(nums) == 1 {
		return nums
	}

	leftSortPart := mergeSort(nums[:len(nums)/2])
	rightSortPart := mergeSort(nums[len(nums)/2:])

	var result []int

	lIndex := 0
	rIndex := 0

	for lIndex < len(leftSortPart) && rIndex < len(rightSortPart) {
		if leftSortPart[lIndex] < rightSortPart[rIndex] {
			result = append(result, leftSortPart[lIndex])
			lIndex++
		} else {
			result = append(result, rightSortPart[rIndex])
			rIndex++
		}
	}

	if lIndex == len(leftSortPart) {
		return append(result, rightSortPart[rIndex:]...)
	} else if rIndex == len(rightSortPart) {
		return append(result, leftSortPart[lIndex:]...)
	}
	return result
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
