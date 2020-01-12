package main

func rob2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	curMaxRob := 0
	preMaxRob := 0
	for _, number := range nums {
		tmp := curMaxRob
		curMaxRob = getTwoMaxNumber(preMaxRob+number, curMaxRob)
		preMaxRob = tmp
	}
	return curMaxRob
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	robRes := make([]int, len(nums))

	getMaxRob(len(nums)-1, nums, robRes)

	return robRes[len(nums)-1]
}

func getMaxRob(index int, nums []int, robRes []int) bool {
	if index == 0 {
		robRes[0] = nums[0]
		return true
	}

	robPre := getMaxRob(index-1, nums, robRes)

	if robPre {
		if index == 1 {
			if nums[index] > robRes[index-1] {
				robRes[index] = nums[1]
				return true
			} else {
				robRes[index] = robRes[index-1]
				return false
			}
		} else {
			if nums[index]+robRes[index-2] > robRes[index-1] {
				robRes[index] = nums[index] + robRes[index-2]
				return true
			} else {
				robRes[index] = robRes[index-1]
				return false
			}
		}
	} else {
		robRes[index] = robRes[index-1] + nums[index]
		return true
	}

}
