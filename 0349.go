package main

func intersection(nums1 []int, nums2 []int) []int {
	num2Map := make(map[int]int)
	resMap := make(map[int]int)

	var res []int

	for _, i2 := range nums2 {
		num2Map[i2] = 1
	}

	for _, number := range nums1 {
		if num2Map[number] != 0 {
			resMap[number] = 1
		}
	}

	for i := range resMap {
		res = append(res, i)
	}

	return res
}
