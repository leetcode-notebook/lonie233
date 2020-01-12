package main

import (
	"strconv"
)

func reverse(x int) int {


	flag := 1
	if x < 0 {
		flag = -1
	}
	x = x * flag

	str := strconv.Itoa(x)
	numbers := 0
	for i := len(str) - 1; i > -1; i-- {
		newNumber, err := strconv.Atoi(string(str[i]))
		nothingh(err)
		numbers = 10*numbers + newNumber

		if numbers*flag < -2147483648 || numbers*flag > 2147483647 {
			return 0
		}
	}

	return numbers * flag
}
func nothingh(err error) {

}
