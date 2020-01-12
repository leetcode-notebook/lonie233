package main

func maxAreaV2(height []int) int {
	left := 0
	right := len(height) - 1

	maxSize := getTwoMinNumber(height[left], height[right]) * right
	for left < right {
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
		tmpMaxSize := getTwoMinNumber(height[left], height[right]) * (right - left)
		if tmpMaxSize > maxSize {
			maxSize = tmpMaxSize
		}
	}

	return maxSize
}

func maxArea(height []int) int {

	maxSize := getTwoMinNumber(height[0], height[1])
	for index := 1; index < len(height); index++ {
		blockMaxSize := getMaxSize(height, index)
		if blockMaxSize > maxSize {
			maxSize = blockMaxSize
		}
	}
	return maxSize
}

func getMaxSize(height []int, index int) int {
	blockMaxSize := getTwoMinNumber(height[index], height[0]) * index
	for i := 1; i < index; i++ {
		tmpSize := getTwoMinNumber(height[index], height[i]) * (index - i)
		if tmpSize > blockMaxSize {
			blockMaxSize = tmpSize
		}
	}
	return blockMaxSize
}
