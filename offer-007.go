package main

//前序遍历 preorder = [3,9,20,15,7]
//中序遍历 inorder = [9,3,15,20,7]
func offerbuildTree(preorder []int, inorder []int) *TreeNode {

	if preorder == nil || inorder == nil {
		return nil
	}

	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	var root = &TreeNode{preorder[0], nil, nil}
	if len(preorder) == 1 || len(inorder) == 1 {
		return root
	}

	childLength := getChildLength(preorder[0], inorder)

	root.Left = offerbuildTree(preorder[1:childLength+1], inorder[:childLength])
	root.Right = offerbuildTree(preorder[childLength+1:], inorder[childLength+1:])

	return root
}

func getChildLength(rootVal int, inorder []int) int {
	for i := range inorder {
		if rootVal == inorder[i] {
			return i
		}
	}

	return 0
}
