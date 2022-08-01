/*
 * @lc app=leetcode id=297 lang=golang
 *
 * [297] Serialize and Deserialize Binary Tree
 */
package main

import (
	utils "reply2future.com/utils"
	"strconv"
	"strings"
)

// [0, 10000] nodes
type TreeNode struct {
	// [-1000, 1000]
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

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var sb strings.Builder
	var traverse func(floor []*TreeNode)
	traverse = func(floor []*TreeNode) {
		if len(floor) == 0 {
			return
		}

		var nextFloor []*TreeNode
		for i, node := range floor {
			if i != 0 {
				sb.WriteByte(',')
			}
			if node != nil {
				sb.WriteString(strconv.Itoa(node.Val))
				nextFloor = append(nextFloor, node.Left, node.Right)
			}
		}
		sb.WriteByte('|')
		traverse(nextFloor)
	}
	traverse([]*TreeNode{root})
	return sb.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	floors := strings.Split(data, "|")
	fLen := len(floors)
	if fLen == 0 || floors[0] == "" {
		return nil
	}
	
	var children []interface{}
	for i := fLen - 1; i >= 0; i-- {
		nodes := strings.Split(floors[i], ",")
		ci := 0
		cLen := len(children)
		var currentNodes []interface{}
		
		for _, v := range nodes {
			if v == "" {
				currentNodes = append(currentNodes, nil)
			} else {
				vc, e := strconv.Atoi(v)
				if e != nil {
					panic(e)
				}
				
				n := &TreeNode{Val: vc}
				if cLen > ci {
					if children[ci] != nil {
						n.Left = children[ci].(*TreeNode)
					}
					if children[ci+1] != nil {
						n.Right = children[ci+1].(*TreeNode)
					}
					ci += 2
				}
				
				currentNodes = append(currentNodes, n)
			}
		}
		children = currentNodes
	}
	return children[0].(*TreeNode)
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
// @lc code=end
func main() {
	ser := Constructor()
	deser := Constructor()

	n1 := &TreeNode{Val: 1}
	n2 := &TreeNode{Val: 2}
	n3 := &TreeNode{Val: 3}
	n4 := &TreeNode{Val: 4}
	n5 := &TreeNode{Val: 5}
	n1.Left = n2
	n1.Right = n3
	n3.Left = n4
	n3.Right = n5
	data1 := ser.serialize(n1)
	utils.AssertEqual1[*TreeNode]("example 1", n1, deser.deserialize, data1)
	utils.AssertEqual1[*TreeNode]("example 2",nil,deser.deserialize,"")
}
