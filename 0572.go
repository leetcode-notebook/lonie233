package main

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if t == nil {
		return true
	}

	if s == nil {
		return false
	}

	return checkTreeEqule(s, t) || isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

//以s，t 为根节点, s 是否包含t
func checkTreeEqule(s *TreeNode, t *TreeNode) bool {

	if t == nil && s == nil {
		return true
	}

	if s == nil || t == nil {
		return false
	}

	return s.Val == t.Val && checkTreeEqule(s.Left, t.Left) && checkTreeEqule(s.Right, t.Right)

}
