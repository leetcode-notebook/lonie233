package main

func inorderTraversal(root *TreeNode) []int {
	//if root == nil {
	//	return nil
	//}
	//
	//resLeft := append(inorderTraversal(root.Left))
	//resLeft = append(resLeft, root.Val)
	//var res = append(resLeft, inorderTraversal(root.Right)...)
	//return res
	return inorderTraversalWithLoop(root)
}

func inorderTraversalWithLoop(root *TreeNode) []int {
	var res []int
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return append(res, root.Val)
	}

	var nodeArray []*TreeNode

	nodeArray = append(nodeArray, root)

	for len(nodeArray) > 0 {
		curNode := nodeArray[len(nodeArray)-1]
		if curNode.Left != nil {
			nodeArray = append(nodeArray, curNode.Left)
		} else {
			//中序节点弹出
			res = append(res, curNode.Val)
			nodeArray = nodeArray[:len(nodeArray)-1]
			if len(nodeArray)-1 > -1 {
				nodeArray[len(nodeArray)-1].Left = nil
			}
			if curNode.Right != nil {
				nodeArray = append(nodeArray, curNode.Right)
			}
		}
	}

	return res

}
