/*
 * @lc app=leetcode id=501 lang=golang
 *
 * [501] Find Mode in Binary Search Tree
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
func findMode(root *TreeNode) []int {
    return bruteForce501(root)
}

// #1 brute force =================================================================
func bruteForce501(root *TreeNode) []int {
	memo := make(map[int]int)
	cMax := 0
	cMaxList := make([]int, 0)
	var travelFn func (*TreeNode)
	travelFn = func (root *TreeNode) {
		if root == nil {
			return
		}
		
		memo[root.Val]++
		if memo[root.Val] > cMax {
			cMaxList = []int{root.Val}
			cMax = memo[root.Val]
		} else if memo[root.Val] == cMax {
			cMaxList = append(cMaxList, root.Val)
		}
		
		travelFn(root.Left)
		travelFn(root.Right)	
	}
	travelFn(root)
	
	return cMaxList
}
// #1 brute force =================================================================

// @lc code=end
func main() {
	n1 := &TreeNode{Val: 1}
	n2 := &TreeNode{Val: 2}
	n3 := &TreeNode{Val: 2}
	n1.Right = n2
	n2.Left = n3
	utils.AssertEqual1[[]int,*TreeNode]("example 1",[]int{2},findMode,n1)
	n4 := &TreeNode{Val: 0}
	utils.AssertEqual1[[]int,*TreeNode]("example 2",[]int{0},findMode,n4)
}
