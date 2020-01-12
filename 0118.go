package main

func generate(numRows int) [][]int {
	res := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		subArray := make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				subArray[j] = 1
			} else {
				subArray[j] = res[i-1][j-1] + res[i-1][j]
			}
		}
		res[i] = subArray
	}

	return res
}
