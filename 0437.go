package main

func pathSum(root *TreeNode, sum int) int {

	if root == nil {
		return 0
	}

	return calCurrentIs(root,sum) + pathSum(root.Left,sum)+ pathSum(root.Right,sum)
}

//判断当前节点是否能满足sum
func calCurrentIs(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}

	diff := sum - root.Val
	var diffCount = 0
	if diff == 0 {
		diffCount = 1
	}
	 diffCount += calCurrentIs(root.Left, diff) + calCurrentIs(root.Right, diff)


	return diffCount
}

//获取以当前节点为根的左右子树内有多少满足条件的路径
