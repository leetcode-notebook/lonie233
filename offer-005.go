package main

func replaceSpace(s string) string {

	if len(s) == 0 {
		return s
	}

	result := ""

	for i := 0; i < len(s); i++ {
		tmp := string(s[i])

		if tmp == " " {
			tmp = "%20"
		}
		result += tmp
	}

	return result
}
