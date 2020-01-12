package main

func hasCycle(head *ListNode) bool {
	var slowNode = head
	var fastNode  = head

	for slowNode != nil  && fastNode.Next != nil && fastNode.Next.Next != nil {
		slowNode = slowNode.Next
		fastNode = fastNode.Next.Next
		if fastNode == slowNode{
			return true
		}
	}
	return false
}
