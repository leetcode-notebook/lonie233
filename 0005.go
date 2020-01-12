package main

func longestPalindrome(s string) string {

	res := ""
	var resMap [][]bool
	curMax := 0

	for i := 0; i < len(s); i++ {
		var tmpArr = make([]bool, len(s))
		for j := 0; j < len(s); j++ {
			tmpArr[j] = false
		}
		resMap = append(resMap, tmpArr)
	}

	for length := 1; length <= len(s); length++ {
		for start := 0; start < len(s); start++ {
			end := start + length - 1
			if end >= len(s) {
				break
			}
			resMap[start][end] = ((length == 1 || length == 2) || resMap[start+1][end-1]) && s[start] == s[end]
			if resMap[start][end] && length > curMax {
				res = s[start : end+1]
			}
		}
	}

	return res
}
