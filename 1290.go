package main

import "math"

func getDecimalValue(head *ListNode) int {
	tmp := head
	length := 0
	res := 0
	for tmp != nil {
		length++
		tmp = tmp.Next
	}

	for head != nil {
		if head.Val == 1 {
			res += (int)(math.Pow(2, float64(length-1)))
		}
		length--
		head = head.Next
	}
	return res
}
