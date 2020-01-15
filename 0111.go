package main

func minDepth(root *TreeNode) int {
	if root == nil {
		 return 0
	}

	return 1 + getTwoMinNumber(minDepth(root.Left), minDepth(root.Right))
}
