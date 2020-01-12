package main

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil {
		return head
	}

	var pre = &ListNode{-1,head}
	var tmpPre = pre
	var mNode *ListNode = nil
	var nNode *ListNode = nil

	for  n > -1 {
		if m == 1 {
			mNode = pre
		}

		if n == 0 {
			nNode = pre
		}
		m--
		n--
		pre = pre.Next
	}

	var lastNode = nNode.Next

	nNode.Next = nil

	var readLast = mNode.Next
	//拼接翻转头
	mNode.Next = reverseRecursion(mNode.Next)
	//拼接翻转尾
	readLast.Next = lastNode

	return tmpPre.Next
}

func reverseRecursion(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var newHead = reverseRecursion(head.Next)
	head.Next.Next = head
	head.Next = nil

	return newHead
}
