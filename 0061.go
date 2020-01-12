package main

func rotateRight(head *ListNode, k int) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	var tmpHead = head
	var length = 0

	for head.Next != nil {
		length++
		head = head.Next
	}
	length++

	if length == k {
		return tmpHead
	}

	head.Next = tmpHead

	var index = k % length
	var dis = length - index
	var newHead = tmpHead
	var start = false
	for length-1 > 0 {
		if dis == 0 {
			newHead = tmpHead
			start = true
		}
		dis--
		tmpHead = tmpHead.Next
		if start {
			length--
		}
	}

	tmpHead.Next = nil

	return newHead
}
