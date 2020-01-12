package main

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return head
	}

	var pre = &ListNode{-1, head}
	var newHead = pre

	var lessHead = &ListNode{-1, nil}
	var realLessHead = lessHead

	for head != nil {

		if head.Val < x {
			var newNode = &ListNode{head.Val, nil}
			//拼接到新列表里
			lessHead.Next = newNode
			lessHead = lessHead.Next
			//从老列表里删除
			pre.Next = head.Next
		} else {
			pre = pre.Next
		}
		head = head.Next
	}

	lessHead.Next = newHead.Next

	return realLessHead.Next
	//[1,4,3,2,5,2]
	//3
}
