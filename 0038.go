package main

import "strconv"

func countAndSay(n int) string {
	resMap := make([]string, n)

	resMap[0]  = "1"

	for column := 1; column < n; column++ {
		resMap[column] = parseStr(resMap[column-1])
	}

	return resMap[n-1]
}

func parseStr(lastCol string) string {

	res := ""
	start := 0
	curChar := string(lastCol[0])
	count := 0
	for start < len(lastCol) {
		if curChar == string(lastCol[start]) {
			count++
		} else {
			res += strconv.Itoa(count) + curChar
			curChar = string(lastCol[start])
			count = 1
		}
		start ++
	}
	res += strconv.Itoa(count) + curChar
	return res
}
