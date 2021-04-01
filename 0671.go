package main

func findSecondMinimumValue(root *TreeNode) int {
	if root.Left == nil && root.Right == nil {
		return -1
	}

	leftVal := root.Left.Val
	rightVal := root.Right.Val

	if root.Val == leftVal {
		leftVal = findSecondMinimumValue(root.Left)
	}

	if rightVal == root.Val {
		rightVal = findSecondMinimumValue(root.Right)
	}

	if leftVal == -1 && rightVal == -1 {
		return getTwoMinNumber(leftVal, rightVal)
	}

	if leftVal == -1 {
		return rightVal
	}

	if rightVal == -1 {
		return leftVal
	}
	return 0

}

//第二大
//func findSecondMinimumValue(root *TreeNode) int {
//	leafNodes := getLeafNode(root)
//
//	var maxNodeVal = 0
//	var secNodeVal = 0
//
//	for _, node := range leafNodes {
//		if node.Val > maxNodeVal {
//			secNodeVal = maxNodeVal
//			maxNodeVal = node.Val
//		} else {
//			if node.Val > secNodeVal {
//				secNodeVal = node.Val
//			}
//		}
//	}
//	return secNodeVal
//}
//
//func getLeafNode(root *TreeNode) []*TreeNode {
//
//	var leafList []*TreeNode
//
//	if root == nil {
//		return nil
//	}
//
//	if root.Left == nil && root.Right == nil {
//		leafList = append(leafList, root)
//		return leafList
//	}
//
//	leafList = append(leafList, getLeafNode(root.Left)...)
//	leafList = append(leafList, getLeafNode(root.Right)...)
//	return leafList
//}
