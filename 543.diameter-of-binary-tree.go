/*
 * @lc app=leetcode id=543 lang=golang
 *
 * [543] Diameter of Binary Tree
 */
package main

import (
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
func max543(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func diameterOfBinaryTree(root *TreeNode) int {
	max := -1
	memo := make(map[*TreeNode]int)

	var depthOfBinaryTree543 func(*TreeNode) int
	depthOfBinaryTree543 = func(root *TreeNode) int {
		if root == nil || (root.Left == nil && root.Right == nil) {
			return 0
		}

		if v, ok := memo[root]; ok {
			return v
		}

		sum := 1 + max543(depthOfBinaryTree543(root.Left), depthOfBinaryTree543(root.Right))
		memo[root] = sum 

		return sum
	}

	var diameter543 func(*TreeNode)
	diameter543 = func(node *TreeNode) {
		if node == nil {
			return
		}
		sum := 0
		if node.Left != nil {
			sum += 1
			sum += depthOfBinaryTree543(node.Left)
		}
		if node.Right != nil {
			sum += 1
			sum += depthOfBinaryTree543(node.Right)
		}

		if max < sum {
			max = sum
		}

		diameter543(node.Left)
		diameter543(node.Right)
	}
	diameter543(root)

	return max
}

// @lc code=end
func main() {
	n1 := &TreeNode{Val: 1}
	n2 := &TreeNode{Val: 2}
	n3 := &TreeNode{Val: 3}
	n4 := &TreeNode{Val: 4}
	n5 := &TreeNode{Val: 5}

	n2.Left = n4
	n2.Right = n5
	n1.Left = n2
	n1.Right = n3
	utils.AssertEqual1[int, *TreeNode]("example 1", 3, diameterOfBinaryTree, n1)
}
