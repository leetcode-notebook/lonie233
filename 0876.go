package main

func middleNode(head *ListNode) *ListNode {
	var slowNode = head
	var fastNode = head

	for fastNode != nil && fastNode.Next != nil   {
		slowNode = slowNode.Next
		fastNode = fastNode.Next.Next
	}
	return slowNode

}
