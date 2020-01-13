package main

func buildTreeByPostOrder(inorder []int, postorder []int) *TreeNode {

	if postorder == nil || len(postorder) == 0 {
		return nil
	}
	var root = &TreeNode{postorder[len(postorder)-1], nil, nil}

	if len(postorder) == 1 || len(inorder) == 1 {
		return root
	}

	curInOrderLength := getCurRootIndexInInOrderByPost(postorder, inorder)
	root.Left = buildTreeByPostOrder(inorder[:curInOrderLength],postorder[:curInOrderLength])
	root.Right = buildTreeByPostOrder(inorder[curInOrderLength+1:],postorder[curInOrderLength:len(postorder)-1])

	return root
}

//返回中序遍历的左右子树长度
func getCurRootIndexInInOrderByPost(postOrder []int, inorder []int) int {
	for index, value := range inorder {
		if value == postOrder[len(postOrder)-1] {
			return index
		}
	}

	return -1
}
