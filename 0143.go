package main

func reorderList(head *ListNode) {

	if head == nil || head.Next == nil {
		return
	}

	var node []*ListNode

	for head != nil {
		node = append(node, head)
		head = head.Next
	}

	var length = len(node)

	for i := 0; i < length/2; i++ {
		var insertNode = node[length-i-1]
		var originNode = node[i]
		insertNode.Next = originNode.Next
		originNode.Next = insertNode
	}
	node[length/2].Next = nil
}
