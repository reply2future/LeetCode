package main

/**
 * https://leetcode.com/explore/learn/card/recursion-i/253/conclusion/2384/
 */

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return _combine(1, n)
}

func _combine(start int, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}
	var _ret []*TreeNode
	for i := start; i <= end; i++ {
		_left := _combine(start, i-1)
		_right := _combine(i+1, end)
		for _, _leftRoot := range _left {
			for _, _rightRoot := range _right {
				_midRoot := &TreeNode{Val: i, Left: _leftRoot, Right: _rightRoot}
				_ret = append(_ret, _midRoot)
			}
		}
	}
	return _ret
}
