package main

func generateTrees(n int) []*TreeNode {

	if n==0 {
		return nil
	}

	return generateBST(1, n)
}

func generateBST(start int, end int) []*TreeNode {
	var res []*TreeNode
	if start > end {
		res = append(res, nil)
		return res
	}

	for i := start; i <= end; i++ {
		resLeft := generateBST(start, i-1)
		resRight := generateBST(i+1, end)

		for _, leftNode := range resLeft {
			for _, rightNode := range resRight {
				var curNode = &TreeNode{i, leftNode, rightNode}
				res = append(res, curNode)
			}
		}
	}
	return res
}
