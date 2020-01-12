package main

func minimumTotal(triangle [][]int) int {
	var res [][]int

	for row, numbers := range triangle {
		var rowRes []int
		for col, number := range numbers {
			if row == 0 {
				rowRes = append(rowRes, number)
			} else if col == 0 {
				rowRes = append(rowRes, number+res[row-1][col])
			} else {
				if col == len(res[row-1]) {
					rowRes = append(rowRes, number+res[row-1][col-1])
				} else {
					rowRes = append(rowRes, number+getTwoMinNumber(res[row-1][col], res[row-1][col-1]))
				}
			}
		}
		res = append(res, rowRes)
	}

	result := res[len(triangle)-1][0]
	for _, number := range res[len(triangle)-1] {
		if number < result {
			result = number
		}
	}
	return result

}
