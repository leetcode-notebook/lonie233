package main

//以root 为根节点的左右同值最长路径
func longestUnivaluePath(root *TreeNode) int {
	if root == nil {
		return 0
	}

	currentLength := getSameValTree(root.Val, root.Left) + getSameValTree(root.Val, root.Right)

	return getTwoMaxNumber(getTwoMaxNumber(currentLength, longestUnivaluePath(root.Left)), longestUnivaluePath(root.Right))

}

func getSameValTree(val int, root *TreeNode) int {
	if root == nil || val != root.Val {
		return 0
	}

	return 1 + getTwoMaxNumber(getSameValTree(val, root.Left), getSameValTree(val, root.Right))
}
