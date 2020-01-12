package main

func getRow(numRows int) []int {
	res := make([][]int, numRows+1)

	for i := 0; i < numRows+1; i++ {
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

	return res[numRows]
}

