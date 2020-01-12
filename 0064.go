package main

func minPathSum(grid [][]int) int {
	resMap := make([][]int, len(grid))

	for i := 0; i < len(grid); i++ {
		tmpArr := make([]int, len(grid[i]))
		resMap[i] = tmpArr
		for j := 0; j < len(grid[i]); j++ {
			if i == 0 && j == 0 {
				resMap[i][j] = grid[i][j]
			} else if i == 0 {
				resMap[i][j] = grid[i][j] + resMap[i][j-1]
			} else if j == 0 {
				resMap[i][j] = grid[i][j] + resMap[i-1][j]
			} else {
				resMap[i][j] = grid[i][j] + getTwoMinNumber(resMap[i-1][j], resMap[i][j-1])
			}
		}
	}
	columns := len(grid) - 1
	return resMap[columns][len(grid[columns])-1]
}
