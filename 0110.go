package main

import "math"

func isBalanced(root *TreeNode) bool {

	if root == nil {
		return true
	}

	leftHeight := getTreeHeight(root.Left)
	rightHeight := getTreeHeight(root.Right)

	return math.Abs(leftHeight-rightHeight) < 2 && isBalanced(root.Left) && isBalanced(root.Right)
}

func getTreeHeight(root *TreeNode) float64 {
	if root == nil {
		return 0
	}

	return math.Max(getTreeHeight(root.Left), getTreeHeight(root.Right)) + 1
}
