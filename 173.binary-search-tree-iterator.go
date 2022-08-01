/*
 * @lc app=leetcode id=173 lang=golang
 *
 * [173] Binary Search Tree Iterator
 */
package main
import (
	utils "reply2future.com/utils"
	"fmt"
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
type BSTIterator struct {
    Cursors []*TreeNode
}

func (this *BSTIterator) AddLeftNodeToStack(root *TreeNode) {
	for root != nil {
		this.Cursors = append(this.Cursors, root)
		root = root.Left
	}
}

func Constructor(root *TreeNode) BSTIterator {
	bi := BSTIterator{}
	bi.AddLeftNodeToStack(root)
	return bi
}


func (this *BSTIterator) Next() int {
	cLen := len(this.Cursors)
	node := this.Cursors[cLen - 1]
	this.Cursors = this.Cursors[:cLen - 1]

	this.AddLeftNodeToStack(node.Right)
	return node.Val
}


func (this *BSTIterator) HasNext() bool {
    return len(this.Cursors) > 0
}


/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
// @lc code=end
func main() {
	fmt.Println("BST Iterator 1 ------------")
	r1 := &TreeNode{Val: 7}
	r2 := &TreeNode{Val: 3}
	r3 := &TreeNode{Val: 15}
	r4 := &TreeNode{Val: 9}
	r5 := &TreeNode{Val: 20}
	r1.Left = r2
	r1.Right = r3
	r3.Left = r4
	r3.Right = r5
	b1 := Constructor(r1)

	utils.AssertEqual0[int]("next()",3,b1.Next)
	utils.AssertEqual0[int]("next()",7,b1.Next)
	utils.AssertEqual0[bool]("hasNext()",true,b1.HasNext)
	utils.AssertEqual0[int]("next()",9,b1.Next)
	utils.AssertEqual0[bool]("hasNext()",true,b1.HasNext)
	utils.AssertEqual0[int]("next()",15,b1.Next)
	utils.AssertEqual0[bool]("hasNext()",true,b1.HasNext)
	utils.AssertEqual0[int]("next()",20,b1.Next)
	utils.AssertEqual0[bool]("hasNext()",false,b1.HasNext)
}
