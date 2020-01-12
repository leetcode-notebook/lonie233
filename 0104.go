package main

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + getTwoMaxNumber(maxDepth(root.Left), maxDepth(root.Right))
}
