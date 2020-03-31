package main

func reversePrint(head *ListNode) []int {

	length := 0

	countHead := head

	for countHead != nil {
		length++
		countHead = countHead.Next
	}

	var nodeArray = make([]int, length)

	for i := length - 1; i > -1; i-- {
		nodeArray[i] = head.Val
		head = head.Next
	}

	return nodeArray
}
