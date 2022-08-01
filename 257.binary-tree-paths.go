/*
 * @lc app=leetcode id=257 lang=golang
 *
 * [257] Binary Tree Paths
 */
package main

import (
	utils "reply2future.com/utils"
	"strconv"
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
func binaryTreePaths(root *TreeNode) []string {
	if isLeaf(root) {
		return []string{strconv.Itoa(root.Val)}
	}

	var ret []string

	if root.Left != nil {
		children := binaryTreePaths(root.Left)
		for _, childPath := range children {
			ret = append(ret, strconv.Itoa(root.Val)+"->"+childPath)
		}
	}

	if root.Right != nil {
		children := binaryTreePaths(root.Right)
		for _, childPath := range children {
			ret = append(ret, strconv.Itoa(root.Val)+"->"+childPath)
		}
	}

	return ret
}

func isLeaf(root *TreeNode) bool {
	return root.Left == nil && root.Right == nil
}

// @lc code=end
func main() {
	n1 := &TreeNode{Val: 1}
	n2 := &TreeNode{Val: 2}
	n3 := &TreeNode{Val: 3}
	n4 := &TreeNode{Val: 5}

	n1.Left = n2
	n1.Right = n3
	n2.Right = n4

	utils.AssertEqual1[[]string, *TreeNode]("example 1", []string{"1->2->5", "1->3"}, binaryTreePaths, n1)

	n5 := &TreeNode{Val: 1}
	utils.AssertEqual1[[]string, *TreeNode]("example 2", []string{"1"}, binaryTreePaths, n5)
}
