package main

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	_leftMax := maxDepth(root.Left)
	_rightMax := maxDepth(root.Right)
	if _leftMax > _rightMax {
		return _leftMax + 1
	}
	return _rightMax + 1
}
