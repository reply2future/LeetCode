/*
 * @lc app=leetcode id=530 lang=golang
 *
 * [530] Minimum Absolute Difference in BST
 */
package main

import (
	"math"
	utils "reply2future.com/utils"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getMinimumDifference(root *TreeNode) int {
	min := math.MaxInt

	var travel func(tree *TreeNode)
	travel = func(tree *TreeNode) {
		if tree == nil {
			return
		}

		if tree.Left != nil {
			c := tree.Left
			for c.Right != nil {
				c = c.Right
			}
			if tree.Val - c.Val < min {
				min = tree.Val - c.Val
			}
		}

		if tree.Right != nil {
			c := tree.Right
			for c.Left != nil {
				c = c.Left
			}
			if c.Val - tree.Val < min {
				min = c.Val - tree.Val
			}
		}

		travel(tree.Left)
		travel(tree.Right)
	}
	travel(root)

	return min
}

// @lc code=end
func main() {
	n1 := &TreeNode{Val: 1}
	n2 := &TreeNode{Val: 2}
	n3 := &TreeNode{Val: 3}
	n4 := &TreeNode{Val: 4}
	n5 := &TreeNode{Val: 6}
	n2.Left = n1
	n2.Right = n3
	n4.Left = n2
	n4.Right = n5

	utils.AssertEqual1[int, *TreeNode]("example 1", 1, getMinimumDifference, n4)
}
