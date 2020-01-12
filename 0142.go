package main

func detectCycle2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	var spec = getSpecNode(head)

	if spec == nil {
		return nil
	}

	for head != nil {
		if head == spec {
			return head
		}
		head = head.Next
		spec = spec.Next
	}

	return nil
}

func getSpecNode(head *ListNode) *ListNode {
	var slow = head
	var fast = head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return slow
		}
	}
	return nil
}
