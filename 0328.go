package main

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}

	nextOddNode := oddEvenList(head.Next.Next)

	head.Next = addMidNode(nextOddNode, head.Next)
	return head
}

func addMidNode(head *ListNode, midNode *ListNode) *ListNode {

	slow := head
	fast := head

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	midNode.Next = slow.Next
	slow.Next = midNode

	return head
}

func oddEvenList2(head *ListNode) *ListNode {

	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}

	odd := head
	even := head.Next
	evenHead := even

	for even != nil && even.Next != nil  {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}

	odd.Next = evenHead

	return head
}
