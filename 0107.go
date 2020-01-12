package main

func levelOrderBottom(root *TreeNode) [][]int {

	if root == nil {
		return nil
	}

	//构造当前节点的左右子节点
	var curLevel []int
	curLevel = append(curLevel, root.Val)

	resLeft := levelOrderBottom(root.Left)
	resRight := levelOrderBottom(root.Right)

	var res [][]int
	var rightBigLeft = false
	//拼接左右子节点的构建结果
	if len(resRight) > len(resLeft) {
		rightBigLeft = true
		tmp := resLeft
		resLeft = resRight
		resRight = tmp
	}

	diff := len(resLeft) - len(resRight)

	for i := 0; i < len(resLeft); i++ {
		if i-diff > -1 {
			if rightBigLeft {
				resLeft[i] = append(resRight[i-diff], resLeft[i]...)
			} else {
				resLeft[i] = append(resLeft[i], resRight[i-diff]...)
			}
		}
		res = append(res, resLeft[i])
	}

	//当前节点的左右子节点拼接到最后
	res = append(res, curLevel)
	return res

}
