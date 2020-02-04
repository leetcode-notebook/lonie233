package main

import "math"

func diameterOfBinaryTree(root *TreeNode) int {

	if root == nil{
		 return 0
	}

	return int(getMaxDistance(root, 0))
}

func getMaxDistance(root *TreeNode, distance float64) float64 {

	leftHeight := getTreeHeight(root.Left)
	rightHeight := getTreeHeight(root.Right)

	curMaxDis := leftHeight + rightHeight - 1

	if curMaxDis > distance {
		distance = curMaxDis
	}

	leftSubMaxDis := getMaxDistance(root.Left, distance)
	rightSubMaxDis := getMaxDistance(root.Right, distance)

	return math.Max(leftSubMaxDis,rightSubMaxDis);

}
