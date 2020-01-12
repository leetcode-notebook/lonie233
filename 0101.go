package main

import "fmt"

func isSymmetric(root *TreeNode) bool {

	if root == nil {
		return true
	}

	var midResArr []int
	var nodeArr []*TreeNode
	nodeArr = append(nodeArr, root)

	for len(nodeArr) > 0 {
		curNode := nodeArr[len(nodeArr)-1]
		if curNode.Left != nil || curNode.Right != nil {
			midResArr = append(midResArr, curNode.Val)
			nodeArr = append(nodeArr, curNode.Right, curNode.Left)
			midResArr = append(midResArr, curNode.Left.Val, curNode.Right.Val)
		} else {
			nodeArr = nodeArr[:len(nodeArr)-1]
			if len(nodeArr)-1 > -1 {
				nodeArr[len(nodeArr)-1].Left = nil
				nodeArr[len(nodeArr)-1].Right = nil
			}
		}
	}

	fmt.Println(midResArr)

	return true
}
