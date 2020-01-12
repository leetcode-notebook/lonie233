package main

func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var newHead = &ListNode{head.Val, nil}
	newHead.Next = nil

	head = head.Next
	for head != nil {
		var sortNode = head
		head = head.Next
		sortNode.Next = nil
		newHead = addSortNode(newHead, sortNode)
	}

	return newHead
}

func addSortNode(head *ListNode, sortNode *ListNode) *ListNode {
	var pre = &ListNode{-1, head}
	var newNode = pre
	for head != nil {
		if head.Val > sortNode.Val {
			pre.Next = sortNode
			sortNode.Next = head
			return newNode.Next
		}
		head = head.Next
		pre = pre.Next
	}
	pre.Next = sortNode
	return newNode.Next
}
