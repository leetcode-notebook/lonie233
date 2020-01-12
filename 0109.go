package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {

	if head == nil {
		return nil
	}

	if head.Next == nil {
		return &TreeNode{head.Val, nil, nil}
	}

	var preNode = &ListNode{-1, head}
	var slowNode = head
	var fastNode = head

	for fastNode != nil && fastNode.Next != nil {
		preNode = preNode.Next
		slowNode = slowNode.Next
		fastNode = fastNode.Next.Next
	}

	var root = &TreeNode{slowNode.Val, nil, nil}
	preNode.Next = nil
	root.Left = sortedListToBST(head)
	root.Right = sortedListToBST(slowNode.Next)

	return root
}
