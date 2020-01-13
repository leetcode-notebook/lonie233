package main

func brokenCalc(X int, Y int) int {

	count := 0

	for X < Y {
		count++
		if Y%2 == 0 {
			Y = Y / 2
		} else {
			Y = Y+1
		}

	}
	return count + X - Y

}
