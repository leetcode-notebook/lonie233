package main

func numDecodings(s string) int {
	if len(s) == 0 {
		return 0
	}
	length := len(s)
	res := 0
	help := 1

	if s[length-1] != '0' {
		res = 1
	}

	for i := length - 2; i > -1; i-- {
		if s[i] == '0' {
			help = res
			res = 0
			continue
		}

		if (s[i]-'0')*10+(s[i+1]-'0') < 27 {
			res += help
			help = res - help
		} else {
			help = res
		}

	}
	return res
}
