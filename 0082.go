package main

func deleteDuplicates2(head *ListNode) *ListNode {

	if head.Next == nil {
		return head
	}

	if head.Val == head.Next.Val {
		for head.Next != nil && head.Val == head.Next.Val {
			head = head.Next
		}
		return deleteDuplicates2(head.Next)
	} else {
		head.Next = deleteDuplicates2(head.Next)
	}

	return head
}
