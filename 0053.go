package main

func maxSubArray(nums []int) int {

	//if len(nums) == 0 {
	//	return 0
	//}
	//maxSum := sumMaxXSubArray(nums, 0)
	//
	//for index := 1; index < len(nums); index++ {
	//	xSum := sumMaxXSubArray(nums, index)
	//	if xSum > maxSum {
	//		maxSum = xSum
	//	}
	//}
	//
	//return maxSum
	var res  = make([]int, len(nums))
	sumReclusion(res, nums, 0)

	result := res[0]
	for _,e := range res {
		if e > result {
			result = e
		}
	}

	return result
}

func sumReclusion(res []int, nums []int, index int) int {
	if index == len(nums)-1 {
		res[index] = nums[index]
		return nums[index]
	}

	nextMaxSum := sumReclusion(res, nums, index+1)
	newMaxSum := nums[index]
	if nextMaxSum > 0 {
		newMaxSum = nums[index] + nextMaxSum
	}

	res[index] = newMaxSum
	return newMaxSum
}

func sumMaxXSubArray(nums []int, index int) int {
	maxSum := nums[index]
	tmpSum := maxSum
	for i := index + 1; i < len(nums); i++ {
		tmpSum += nums[i]
		if tmpSum > maxSum {
			maxSum = tmpSum
		}
	}
	return maxSum
}
