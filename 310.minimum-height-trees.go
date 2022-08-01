/*
 * @lc app=leetcode id=310 lang=golang
 *
 * [310] Minimum Height Trees
 */
package main

import (
	utils "reply2future.com/utils"
)

// @lc code=start
func findMinHeightTrees(n int, edges [][]int) []int {
	// return bruteForceFindMinHeightTrees(n, edges)
	return topologicalFindMinHeightTrees(n, edges)
}

// #1 brute force =================================
// x time limit
type MhtNode struct {
	Value int
	Next  []*MhtNode
}

type MhtKey struct {
	cNode *MhtNode
	pNode *MhtNode
}

func generateMhtNodes(n int) map[int]*MhtNode {
	ret := make(map[int]*MhtNode)
	for i := 0; i < n; i++ {
		ret[i] = &MhtNode{Value: i}
	}
	return ret
}

func bruteForceFindMinHeightTrees(n int, edges [][]int) []int {
	nodeMaps := generateMhtNodes(n)
	for _, edge := range edges {
		n1 := nodeMaps[edge[0]]
		n2 := nodeMaps[edge[1]]
		n1.Next = append(n1.Next, n2)
		n2.Next = append(n2.Next, n1)
	}

	memo := make(map[MhtKey]int)
	var findMh func(*MhtNode, *MhtNode) int
	findMh = func(node *MhtNode, parent *MhtNode) int {
		if node == nil || len(node.Next) == 0 {
			return 0
		}

		var restHeight int
		for _, n := range node.Next {
			if parent == n {
				continue
			}
			if v, ok := memo[MhtKey{cNode: n, pNode: node}]; ok {
				if v > restHeight {
					restHeight = v
				}
			} else {
				result := findMh(n, node)
				if result > restHeight {
					restHeight = result
				}
			}
		}
		return 1 + restHeight
	}

	var ret []int
	var height int
	for i := 0; i < n; i++ {
		result := findMh(nodeMaps[i], nodeMaps[i])
		if height == 0 || result == height {
			height = result
			ret = append(ret, i)
		} else if result > height {
			continue
		} else {
			ret = []int{i}
			height = result
		}
	}
	return ret
}

// #1 =============================================

// #2 topological sorting =============================================
func topologicalFindMinHeightTrees(n int, edges [][]int) []int {
	if n <= 1 {
		return []int{0}
	}

	graph := make(map[int][]int)
	inDegrees := make([]int, n)

	for _, edge := range edges {
		if v1, ok1 := graph[edge[0]]; ok1 {
			graph[edge[0]] = append(v1, edge[1])
		} else {
			graph[edge[0]] = []int{edge[1]}
		}

		if v2, ok2 := graph[edge[1]]; ok2 {
			graph[edge[1]] = append(v2, edge[0])
		} else {
			graph[edge[1]] = []int{edge[0]}
		}

		inDegrees[edge[0]]++
		inDegrees[edge[1]]++
	}

	var leafQueue []int

	for i, v := range inDegrees {
		if v == 1 {
			leafQueue = append(leafQueue, i)
		}
	}

	for n > 2 {
		size := len(leafQueue)
		n -= size

		for i := 0; i < size; i++ {
			for _, to := range graph[leafQueue[i]] {
				inDegrees[to]--
				if inDegrees[to] == 1 {
					leafQueue = append(leafQueue, to)
				}
			}
			inDegrees[leafQueue[i]]--
		}

		leafQueue = leafQueue[size:]
	}

	return leafQueue
}
// #2 =================================================================


// @lc code=end
func main() {
	utils.AssertEqual2[[]int, int, [][]int]("example 1", []int{1}, findMinHeightTrees, 4, [][]int{{1, 0}, {1, 2}, {1, 3}})
	utils.AssertEqual2[[]int, int, [][]int]("example 2", []int{3, 4}, findMinHeightTrees, 6, [][]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}})
}
