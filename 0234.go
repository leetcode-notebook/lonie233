package main

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}

	if head.Next == nil {
		return true
	}
	var headA = head
	var fastNode = head
	for head.Next != nil && fastNode != nil {
		head = head.Next
		if fastNode.Next != nil {
			fastNode = fastNode.Next.Next
		} else {
			fastNode = nil
		}
	}

	var lastNode = reverlist(head)

	for lastNode != nil {
		if headA.Val != lastNode.Val {
			return false
		}
		headA = headA.Next
		lastNode = lastNode.Next
	}

	return true
}

func reverlist(head *ListNode) *ListNode {

	if head.Next == nil {
		return head
	}
	var newHead = reverlist(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
