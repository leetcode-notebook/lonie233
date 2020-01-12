package main

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	//if l1.Val == 0 {
	//	return l2
	//}
	//
	//if l2.Val == 0 {
	//	return l1
	//}

	var head = &ListNode{0, nil}
	var tmp = head
	var n = 0
	for l1 != nil || l2 != nil {
		var l1Val = 0
		if l1 != nil {
			l1Val = l1.Val
			l1 = l1.Next
		}

		var l2Val = 0
		if l2 != nil {
			l2Val = l2.Val
			l2 = l2.Next
		}

		var tmpSum = l1Val + l2Val + n
		if tmpSum > 9 {
			n = 1
			tmpSum -= 10
		} else {
			n = 0
		}

		var newNode = &ListNode{tmpSum, nil}
		tmp.Next = newNode
		tmp = tmp.Next
	}

	if n == 1 {
		var newNode = &ListNode{1, nil}
		tmp.Next = newNode
	}

	return head.Next
}

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwoNumberRecursion(l1, l2, 0)
}

func addTwoNumberRecursion(l1 *ListNode, l2 *ListNode, carry int) *ListNode {
	if l1 == nil && l2 == nil {
		if carry == 0 {
			return nil
		} else {
			return &ListNode{carry, nil}
		}
	} else if l1 == nil {
		l1 = &ListNode{0, nil}
	} else if l2 == nil {
		l2 = &ListNode{0, nil}
	}

	var tmpSum = l1.Val + l2.Val + carry
	carry = 0
	if tmpSum > 9 {
		tmpSum = tmpSum - 10
		carry = 1
	}

	l1.Val = tmpSum

	l1.Next = addTwoNumberRecursion(l1.Next, l2.Next, carry)

	return l1
}
