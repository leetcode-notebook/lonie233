package main

func addTwoNumbers3(l1 *ListNode, l2 *ListNode) *ListNode {

	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	len1 := 0
	len2 := 0

	newHead := l1
	newHead2 := l2

	tmp1 := l1
	tmp2 := l2
	for true {
		if tmp1 != nil {
			len1++
			tmp1 = tmp1.Next
		}

		if tmp2 != nil {
			len2++
			tmp2 = tmp2.Next
		}

		if tmp1 == nil && tmp2 == nil {
			break
		}
	}

	if len1 > len2 {
		newHead2 = addNode(len1-len2, l2)
	} else if len2 > len1 {
		newHead = addNode(len2-len1, l1)
	}

	lastCarry := reverseSum(newHead, newHead2)

	if lastCarry != 0 {
		node := &ListNode{lastCarry, newHead}
		return node
	}

	if newHead.Val == 0 && newHead.Next == nil {
		return newHead
	}

	for newHead != nil {
		if newHead.Val != 0 {
			break
		}
		newHead = newHead.Next
	}
	return newHead
}

func addNode(length int, node *ListNode) *ListNode {
	tmpNode := node
	l1 := &ListNode{0, nil}
	newHead := l1
	for i := 0; i < length-1; i++ {
		l1.Next = &ListNode{0, nil}
		l1 = l1.Next
	}
	l1.Next = tmpNode
	return newHead
}

func reverseSum(l1 *ListNode, l2 *ListNode) int {

	if l1 == nil || l2 == nil {
		return 0
	}

	sum := l1.Val + l2.Val + reverseSum(l1.Next, l2.Next)
	l1.Val = sum % 10
	return sum / 10
}
