package main

func uniquePaths(m int, n int) int {
	var res = make([][]int, m)
	for x := 0; x < m; x++ {
		xArr := make([]int, n)
		res[x] = xArr
		for y := 0; y < n; y++ {
			if x == 0 || y == 0 {
				xArr[y] = 1
			} else {
				xArr[y] = res[x-1][y] + res[x][y-1]
			}
		}
	}
	return res[m-1][n-1]
}
