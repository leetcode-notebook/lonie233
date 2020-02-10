package main

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}

	currentLeft := 0
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		currentLeft = root.Left.Val
	}

	return currentLeft + sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
}
