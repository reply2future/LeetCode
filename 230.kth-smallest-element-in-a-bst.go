/*
 * @lc app=leetcode id=230 lang=golang
 *
 * [230] Kth Smallest Element in a BST
 */
package main
import (
	utils "reply2future.com/utils"
)

type TreeNode struct {
    Val int
    Left *TreeNode
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
func kthSmallest(root *TreeNode, k int) int {
	if k < 1 {
		panic("k must be greater than 1")
	}
    // in-order traveral
	var ret int
	var stack []*TreeNode
	cNode := root
	for {
		stack = append(stack, getLeftNodeToStack(cNode)...)
		sLen := len(stack)
		node := stack[sLen-1]
		if k == 1 {
			return node.Val
		}

		stack = stack[:sLen-1]
		cNode = node.Right
		k -= 1
	}
	return ret
}

func getLeftNodeToStack(root *TreeNode) []*TreeNode {
	stack := []*TreeNode{}
	
	for root != nil {
		stack = append(stack, root)
		root = root.Left
	}

	return stack
}
// @lc code=end
func main() {
	n1 := TreeNode{Val:3}
	n2 := TreeNode{Val:1}
	n3 := TreeNode{Val:4}
	n4 := TreeNode{Val:2}
	n1.Left = &n2
	n2.Right = &n3
	n2.Right = &n4
	utils.AssertEqual2[int,*TreeNode,int]("example 1",1,kthSmallest,&n1,1)

	n5 := TreeNode{Val:1}
	n6 := TreeNode{Val:2}
	n7 := TreeNode{Val:3}
	n8 := TreeNode{Val:4}
	n9 := TreeNode{Val:5}
	n10 := TreeNode{Val:6}
	n9.Left = &n7
	n9.Right = &n10
	n7.Right = &n8
	n7.Left = &n6
	n6.Left = &n5
	utils.AssertEqual2[int,*TreeNode,int]("example 2",3,kthSmallest,&n9,3)
}
