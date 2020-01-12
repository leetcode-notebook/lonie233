package main

func swapPairs(head *ListNode) *ListNode {
	if head == nil  {
		return nil
	}

	if head.Next == nil {
		return head
	}

	var nextNode = head.Next.Next

	var tmpNode = head
	head = head.Next
	head.Next = tmpNode

	tmpNode.Next = swapPairs(nextNode)

	return  head
}
