package main

func levelOrder(root *TreeNode) [][]int {

	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		leafRes := make([][]int, 1)
		leafRes[0] = append(leafRes[0], root.Val)
		return leafRes
	}

	resLeft := levelOrder(root.Left)
	resRight := levelOrder(root.Right)

	level := getTwoMaxNumber(len(resRight), len(resLeft))
	curRes := make([][]int, level+1)

	curRes[0] = append(curRes[0], root.Val)

	for i := 0; i < level; i++ {
		if i < len(resLeft) {
			curRes[i+1] = append(curRes[i+1], resLeft[i]...)
		}

		if i < len(resRight) {
			curRes[i+1] = append(curRes[i+1], resRight[i]...)
		}
	}
	return curRes
}


func levelOrderV2(root *TreeNode) [][]int {

	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		leafRes := make([][]int, 1)
		leafRes[0] = append(leafRes[0], root.Val)
		return leafRes
	}

	resLeft := levelOrder(root.Left)
	resRight := levelOrder(root.Right)

	level := getTwoMaxNumber(len(resRight), len(resLeft))
	curRes := make([][]int, level+1)

	curRes[0] = append(curRes[0], root.Val)

	for i := 0; i < level; i++ {
		if i < len(resLeft) {
			curRes[i+1] = append(curRes[i+1], resLeft[i]...)
		}

		if i < len(resRight) {
			curRes[i+1] = append(curRes[i+1], resRight[i]...)
		}
	}
	return curRes
}
