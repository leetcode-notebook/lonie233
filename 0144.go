package main

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var res []int
	res = append(res, root.Val)

	res = append(res, preorderTraversal(root.Left)...)
	res = append(res, preorderTraversal(root.Right)...)

	return res
}

func preorderTraversalWithCycle(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var nodeArr []*TreeNode
	nodeArr = append(nodeArr, root)

	var res []int

	for len(nodeArr) > 0 {
		curNode := nodeArr[len(nodeArr)-1]
		nodeArr = nodeArr[:len(nodeArr)-1]
		//前序遍历值
		res = append(res, curNode.Val)

		if curNode.Right != nil {
			nodeArr = append(nodeArr, curNode.Right)
		}

		if curNode.Left != nil {
			nodeArr = append(nodeArr, curNode.Left)
		}

	}

	return res
}
