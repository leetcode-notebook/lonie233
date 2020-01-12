package main

func rangeSumBST(root *TreeNode, L int, R int) int {
	if L == R {
		return L
	}

	if root == nil {
		return 0
	}

	if root.Val > R {
		return rangeSumBST(root.Left, L, R)
	}
	if root.Val < L {
		return rangeSumBST(root.Right, L, R)
	}

	return root.Val + rangeSumBST(root.Left, L, R) + rangeSumBST(root.Right, L, R)

}
