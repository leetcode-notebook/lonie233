package main

func rob3(root *TreeNode) int {
	if root == nil {
		return 0
	}

	val1 := root.Val

	val2 := 0
	if root.Left != nil {
		val1 += rob3(root.Left.Left) + rob3(root.Left.Right)
	}

	if root.Right != nil {
		val1 += rob3(root.Right.Left) + rob3(root.Right.Right)
	}

	val2 = rob3(root.Left) + rob3(root.Right)
	if val1 > val2 {
		return val1
	} else {
		return val2
	}

}
