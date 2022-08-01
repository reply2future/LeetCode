/*
 * @lc app=leetcode id=572 lang=golang
 *
 * [572] Subtree of Another Tree
 */
package main

import (
	utils "reply2future.com/utils"
)

type TreeNode = utils.TreeNode
// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {

	var compareFn func(*TreeNode, *TreeNode) bool
	compareFn = func(root *TreeNode, subRoot *TreeNode) bool {
		if root.Val != subRoot.Val {
			return false
		}
		if root.Left != nil {
			if subRoot.Left == nil {
				return false
			}
			if !compareFn(root.Left, subRoot.Left) {
				return false
			}
		} else if subRoot.Left != nil {
            return false
        }
		if root.Right != nil {
			if subRoot.Right == nil {
				return false
			}
			if !compareFn(root.Right, subRoot.Right) {
				return false
			}
		} else if subRoot.Right != nil {
            return false
        }
		return true
	}

    found := false
    var travel func (*TreeNode)
    travel = func(node *TreeNode) {
        if found {
            return
        }

        if node == nil {
            return
        }

        if compareFn(node, subRoot) {
            found = true
            return
        }
        
        travel(node.Left)
        travel(node.Right)
    }
    travel(root)

    return found
}

// @lc code=end
func main() {
    a := utils.GenerateTree([]any{3, 4, 5, 1, 2})
    b := utils.GenerateTree([]any{4, 1, 2})
	utils.AssertEqual2[bool, *TreeNode, *TreeNode]("example 1", true, isSubtree, a, b)
}
