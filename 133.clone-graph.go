/*
 * @lc app=leetcode id=133 lang=golang
 *
 * [133] Clone Graph
 */
package main

import (
	// utils "reply2future.com/utils"
	"fmt"
)

// Definition for a Node.
type Node struct {
	Val       int
	Neighbors []*Node
}

// @lc code=start

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	
	// return tileCloneGraph(node)
	return dpCloneGraph(node)
}

// #1 tile map
func tileCloneGraph(node *Node) *Node {
	maps := make(map[int][]int)
	tileNodes(node, maps)

	nm := make(map[int]*Node)
	// value
	for k := range maps {
		nm[k] = &Node{Val: k}
	}

	// relationship
	for k, v := range maps {
		nm[k].Neighbors = []*Node{}
		for _, n := range v {
			nm[k].Neighbors = append(nm[k].Neighbors, nm[n])
		}
	}

	return nm[1]
}

func tileNodes(node *Node, maps map[int][]int) {
	_, ok := maps[node.Val]
	if ok {
		return
	}

	var neighbors []int
	for _, v := range node.Neighbors {
		neighbors = append(neighbors, v.Val)
	}
	maps[node.Val] = neighbors

	for _, v := range node.Neighbors {
		tileNodes(v, maps)
	}
}

func dpCloneGraph(node *Node) *Node {
	memo := make(map[int]*Node)
	var recursionFn func(node *Node) *Node
	recursionFn = func(node *Node) *Node {
		v, ok := memo[node.Val]
		if ok {
			return v
		}

		n := &Node{Val: node.Val, Neighbors: make([]*Node, 0)}
		memo[node.Val] = n
		
		for _, v := range node.Neighbors {
			n.Neighbors = append(n.Neighbors, recursionFn(v))
		}
		
		return n
	}

	return recursionFn(node)
}

// @lc code=end
func main() {
	n1 := &Node{Val: 1}
	n2 := &Node{Val: 2}
	n3 := &Node{Val: 3}
	n4 := &Node{Val: 4}

	n1.Neighbors = []*Node{n2, n4}
	n2.Neighbors = []*Node{n1, n3}
	n3.Neighbors = []*Node{n2, n4}
	n4.Neighbors = []*Node{n1, n3}

	v1 := cloneGraph(n1)
	fmt.Println(v1)

	n5 := &Node{Val: 1}
	v2 := cloneGraph(n5)
	fmt.Println(v2)

	n6 := &Node{Val: 1}
	n7 := &Node{Val: 2}
	n6.Neighbors = []*Node{n7}
	n7.Neighbors = []*Node{n6}
	v3 := cloneGraph(n6)
	fmt.Println(v3)
}
