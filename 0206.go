package main

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var newHead = reverseList(head.Next)

	head.Next = nil
	var tmpHead = newHead
	for tmpHead.Next != nil {
		tmpHead = tmpHead.Next
	}
	tmpHead.Next = head
	return newHead
}

func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var newHead = reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

func reverseList3(head *ListNode) *ListNode {
	var pre *ListNode
	pre = nil

	for head != nil {
		var nextNode = head.Next
		head.Next = pre
		pre = head
		head = nextNode
	}

	return pre
}
