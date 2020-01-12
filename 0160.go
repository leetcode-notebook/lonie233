package main

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var headAA = headA
	var headBB = headB

	for headA != headB {

		if headA == nil {
			headA = headBB
		} else {
			headA = headA.Next
		}

		if headB == nil {
			headB = headAA
		} else {
			headB = headB.Next
		}
	}

	return headA
}
