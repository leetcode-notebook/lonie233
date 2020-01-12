package main



func removeDuplicates02(nums []int) int {

	if len(nums) < 2 {
		return len(nums)
	}

	slow := 0

	for index, value := range nums {
		if nums[slow] != value {
			slow++
			nums[slow] = nums[index]
		}
	}

	return slow + 1
}

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	length := removeDuplicatesRecersion(0, nums)
	return length
}

func removeDuplicatesRecersion(index int, nums []int) int {
	if index == len(nums)-1 {
		return 1
	}

	length := removeDuplicatesRecersion(index+1, nums)

	if nums[index] == nums[index+1] {
		for k, v := range nums {
			if k > index {
				nums[k-1] = v
			}
		}
		return length
	} else {
		return length + 1
	}

}
