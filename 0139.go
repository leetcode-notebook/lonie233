package main

import "strings"

//
//v1: time out
//import "strings"
//
//func wordBreak(s string, wordDict []string) bool {
//	//转移方程 f(i) =  wordDict contain (current word) && f(i+1)
//	//current word = sub(s[0]- s[len-1])
//
//	//1. 直接判断按一个单词处理是否ok
//	if isExist(s, wordDict) {
//		return true
//	}
//
//	//2. move all space
//	s =  strings.Replace(s, " ", "", -1)
//
//	//try find
//	length := len(s)
//
//	for i := 1; i < length; i++ {
//		currWord := s[0:i]
//		if isExist(currWord, wordDict) && wordBreak(s[i:], wordDict) {
//			return true
//		}
//	}
//
//	return false
//}
//
func isExist(word string, wordDict [] string) bool {
	for _, existWord := range wordDict {
		if word == existWord {
			return true
		}
	}
	return false
}

func wordBreak(s string, wordDict []string) bool {

	//2. move all space
	s = strings.Replace(s, " ", "", -1)

	//try find
	length := len(s)
	//dp 表示 s[0,i] 是否符合要求
	dp := make([]bool, length+1)
	dp[0] = true
	for i := 1; i <= length; i++ {
		for j := 0; j <= i; j++ {
			if dp[j] && isExist(s[j:i], wordDict) {
				dp[i] = true
				break
			}
		}
	}
	return dp[length]
}

func checkStr(s string, wordDict []string) bool {
	if isExist(s, wordDict) {
		return true
	}

	length := len(s)

	for i := 0; i < length; i++ {
		if isExist(s[0:i], wordDict) && checkStr(s[i:], wordDict) {
			return true
		}
	}
	return false
}
