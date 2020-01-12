package main

func isSubsequence(s string, t string) bool {

	if s == t || len(s) == 0{
		return true
	}

	for i := len(t) - 1; i > -1; i-- {
		if t[i] == s[len(s)-1] {
			return isSubsequence(s[0:len(s)-1], t[0:i])
		}
	}

	return false
}
