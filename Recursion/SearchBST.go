package main

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == val {
		return root
	}
	_leftR := searchBST(root.Left, val)
	if _leftR != nil {
		return _leftR
	}
	_rightR := searchBST(root.Right, val)
	if _rightR != nil {
		return _rightR
	}
	return nil
}
