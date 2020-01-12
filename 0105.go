package main

//核心思想就是找到根节点，然后递归重建左右子树，难点在于递归划分左右子树的数组
func buildTree(preorder []int, inorder []int) *TreeNode {

	if preorder == nil || len(preorder) == 0 {
		return nil
	}
	var root = &TreeNode{preorder[0], nil, nil}

	if len(preorder) == 1 || len(inorder) == 1 {
		return root
	}

	curInOrderLength := getCurRootIndexInInOrder(preorder, inorder)
	//前序遍历的数组要跨过跟节点
	root.Left = buildTree(preorder[1:curInOrderLength+1], inorder[:curInOrderLength])
	//中序遍历的数据要跨过跟节点
	root.Right = buildTree(preorder[1+curInOrderLength:], inorder[curInOrderLength+1:])

	return root
}

//返回中序遍历的左右子树长度
func getCurRootIndexInInOrder(preorder []int, inorder []int) int {
	for index, value := range inorder {
		if value == preorder[0] {
			return index
		}
	}

	return -1
}
