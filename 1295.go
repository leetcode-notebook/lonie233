package main

func findNumbers(nums []int) int {

	if nums == nil || len(nums) == 0 {
		return 0
	}

	count := 0
	for _, num := range nums {
		count += isEvenLength(num)
	}

	return count

}

//偶数长度返回1，奇数返回0
func isEvenLength(number int) int {

	length := 0
	for number != 0 {
		number = number / 10
		length++
	}

	if length%2 == 0 {
		return 1
	} else {
		return 0
	}
}
