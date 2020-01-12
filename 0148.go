package main

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var newHead = sortList(head.Next)

	if head.Val <= newHead.Val {
		head.Next = newHead
		return head
	} else {
		var temHead = newHead
		var preHead = &ListNode{-1, temHead}
		for temHead != nil {
			if head.Val < temHead.Val {
				head.Next = preHead.Next
				preHead.Next = head
				break
			}
			temHead = temHead.Next
			preHead = preHead.Next
		}
		if temHead == nil {
			preHead.Next = head
			head.Next = nil
		}

		return newHead
	}
}
