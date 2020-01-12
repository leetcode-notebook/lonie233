package main

func isSameTree(p *TreeNode, q *TreeNode) bool {

	if p == nil {
		return q == nil
	}

	if q == nil {
		return false
	}

	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)

}
