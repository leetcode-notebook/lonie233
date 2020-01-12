package main

func removeElements(head *ListNode, val int) *ListNode {

	if head == nil {
		return head
	}

	var pre = &ListNode{-1, head}
	var headA = pre

	for head != nil {
		if head.Val == val {
			pre.Next = head.Next
			head = head.Next
		} else {
			pre = pre.Next
			if pre == nil {
				break
			}
			head = pre.Next
		}
	}
	return headA.Next
}

func removeElements2(head *ListNode, val int) *ListNode {

	if head.Next == nil {
		if head.Val == val {
			return nil
		}
		return head
	}

	head.Next = removeElements2(head.Next, val)

	if head.Val == val {
		return head.Next
	} else {
		return head
	}
}
