package main

func isPalindromeNumber(x int) bool {
	if x < 0 {
		return false
	}

	tmpNumber := x
	reverseNumber := 0

	for tmpNumber != 0 {
		reverseNumber = reverseNumber*10 + tmpNumber%10
		tmpNumber = tmpNumber / 10
	}

	return reverseNumber == x
}

func isPalindromeNumber2(x int) bool {
	if x < 0 {
		return false
	}

	tmpNumber := x
	reverseNumber := 0

	for tmpNumber != 0 {
		reverseNumber = reverseNumber*10 + tmpNumber%10
		tmpNumber = tmpNumber / 10
	}

	return reverseNumber == x
}
