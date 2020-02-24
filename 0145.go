package main


func postorderTraversal(root *TreeNode) []int {

	if root == nil {
		return nil
	}

	var result []int

	result = append(result, postorderTraversal(root.Left)...)
	result = append(result, postorderTraversal(root.Right)...)
	result = append(result, root.Val)
	return result

}

func postorderTraversalByCycle(root *TreeNode) []int {

	if root == nil {
		return nil
	}

	var result []int
	var node []*TreeNode
	node = append(node, root)

	for len(node) > 0 {
		curNode := node[len(node)-1]
		node = node[:len(node)-1]

		var tmpResult []int
		tmpResult = append(tmpResult, curNode.Val)
		tmpResult = append(tmpResult, result...)
		result = tmpResult

		if curNode.Left != nil {
			node = append(node, curNode.Left)
			curNode.Left = nil
		}

		if curNode.Right != nil {
			node = append(node, curNode.Right)
			curNode.Right = nil
		}

	}
	return result
}
