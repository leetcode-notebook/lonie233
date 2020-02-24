package main

func findBottomLeftValue(root *TreeNode) int {
	var nodes []*TreeNode

	nodes = append(nodes, root)

	for nodes != nil {
		cnt := len(nodes)
		for i := 0; i < cnt; i++ {
			if nodes[i].Left != nil {
				nodes = append(nodes, nodes[i].Left)
			}
			if nodes[i].Right != nil {
				nodes = append(nodes, nodes[i].Right)
			}
		}

		if len(nodes) == cnt {
			break
		}

		nodes = nodes[cnt:]
	}
	return nodes[0].Val
}
