package main

func minArray(numbers []int) int {

	if numbers == nil || len(numbers) == 0 {
		return 0
	}

	if len(numbers) == 1 {
		return numbers[0]
	}

	biggerNumber := numbers[0]
	lessNumber := numbers[0]

	for i := 1; i < len(numbers); i++ {
		biggerNumber = numbers[i]
		if lessNumber > biggerNumber {
			return biggerNumber
		}
		lessNumber = biggerNumber

	}

	return numbers[0]

}
