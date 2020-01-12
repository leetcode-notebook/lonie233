package main

func deleteDuplicates(head *ListNode) *ListNode {

	var header *ListNode
	header = &ListNode{0, nil}
	header.Next = head
	var pre = header
	var cur = head

	for cur != nil {

		if cur.Val == pre.Val {
			cur = cur.Next
			pre.Next = cur
		} else {
			pre = pre.Next
			cur = cur.Next
		}
	}

	return header.Next
}

func deleteDuplicatesRecursion(head *ListNode) *ListNode {

	if head == nil ||head.Next == nil {
		return head
	}

	head.Next = deleteDuplicatesRecursion(head.Next)
	if head.Val == head.Next.Val {
		return head.Next.Next
	}

	return head
}
