package main

func sortedArrayToBST(nums []int) *TreeNode {

	if nums == nil || len(nums) == 0 {
		return nil
	}

	if len(nums) == 1 {
		return &TreeNode{nums[0], nil, nil}
	}

	if len(nums) == 2 {
		var root = &TreeNode{nums[1], nil, nil}
		var leftLeaf = &TreeNode{nums[0], nil, nil}
		root.Left = leftLeaf
		return root
	}

	if len(nums) == 3 {
		var root = &TreeNode{nums[1], nil, nil}
		var leftLeaf = &TreeNode{nums[0], nil, nil}
		var rightLeaf = &TreeNode{nums[2], nil, nil}
		root.Left = leftLeaf
		root.Right = rightLeaf
		return root
	}

	curRootIndex := len(nums)/2
	var root = &TreeNode{nums[curRootIndex], nil, nil}
	root.Left = sortedArrayToBST(nums[0:curRootIndex])
	root.Right = sortedArrayToBST(nums[curRootIndex+1:])

	return root
}
