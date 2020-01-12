package main

func splitListToParts(root *ListNode, k int) []*ListNode {
	if root == nil {
		return nil
	}

	length := 0

	tmpHead := root

	for tmpHead != nil {
		length++
		tmpHead = tmpHead.Next
	}

	sliceCount := length / k
	extCount := length % k

	res := make([]*ListNode, k)

	for i := 0; i < k; i++ {
		sliceHead := root
		res = append(res, sliceHead)
		for j := 0; j < sliceCount; j++ {
			root = root.Next
		}

		if extCount > 0 {
			root = root.Next
		}
		sliceTail := root
		root = root.Next
		sliceTail.Next = nil
	}

	return res
}
