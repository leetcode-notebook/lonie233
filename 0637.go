package main

func averageOfLevels(root *TreeNode) []float64 {

	var ret []float64

	var nodes []*TreeNode

	nodes = append(nodes, root)

	for len(nodes) > 0 {
		var cnt = len(nodes)
		var sum = 0

		for i := 0; i < cnt; i++ {
			sum += nodes[cnt].Val
			if nodes[cnt].Left != nil {
				nodes = append(nodes, nodes[cnt].Left)
			}
			if nodes[cnt].Right != nil {
				nodes = append(nodes, nodes[cnt].Right)
			}
		}

		nodes = nodes[cnt:]
		ret = append(ret, float64(sum)/float64(cnt))
	}

	return ret
}
