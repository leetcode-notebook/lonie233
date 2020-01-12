package main

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	var p1Pre *ListNode
	var header *ListNode
	p1Pre = &ListNode{0, nil}
	header = p1Pre

	for l1!= nil && l2 != nil  {
		if l1.Val < l2.Val {
			p1Pre.Next = l1
			l1 = l1.Next
		} else  {
			p1Pre.Next = l2
			l2 = l2.Next
		}
		p1Pre = p1Pre.Next
	}

	if l1==nil {
		 p1Pre.Next = l2
	}

	if l2==nil {
		 p1Pre.Next = l1
	}

	return header.Next
}
