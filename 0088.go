package main

func merge(nums1 []int, m int, nums2 []int, n int) {

	var lastIndex = (m + n) - 1

	for n > 0 && m > 0 {
		curMax := nums1[getMax(m-1)]
		if curMax > nums2[getMax(n-1)] {
			m--
		} else {
			curMax = nums2[getMax(n-1)]
			n--
		}

		nums1[lastIndex] = curMax
		lastIndex--
	}

	for n> 0 && lastIndex > -1 {
		nums1[lastIndex] = nums2[lastIndex]
		lastIndex--
	}

}

func getMax(number int) int {
	if number > 0 {
		return number
	}
	return 0
}
