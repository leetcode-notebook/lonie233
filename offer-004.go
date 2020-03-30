package main

func findNumberIn2DArray(matrix [][]int, target int) bool {

	if nil == matrix || len(matrix) == 0 {
		return false
	}

	rowNumber := 0
	colNumber := len(matrix[0]) - 1

	for colNumber > -1 && rowNumber < len(matrix) {
		if target == matrix[rowNumber][colNumber] {
			return true
		} else if target < matrix[rowNumber][colNumber] {
			colNumber--
		} else {
			rowNumber++
		}
	}

	return false
}
