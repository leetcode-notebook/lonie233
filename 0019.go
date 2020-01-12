package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	if head == nil {
		return head
	}

	var pre = &ListNode{0, head}
	var newHeader = pre
	var endPointer = pre
	n = n + 1

	for n > 0 {
		endPointer = endPointer.Next
		n--
	}

	for endPointer != nil {
		endPointer = endPointer.Next
		pre = pre.Next
	}

	pre.Next = pre.Next.Next

	return newHeader.Next
}
