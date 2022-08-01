/*
 * @lc app=leetcode id=56 lang=golang
 *
 * [56] Merge Intervals
 */
package main

import (
	assert "reply2future.com/utils"
)

// @lc code=start
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func merge(intervals [][]int) [][]int {
	// return mergeBruteForce(intervals)
	return mergeBySingleLinkedTable(intervals)
}

func mergeTwoElements(a []int, b []int) [][]int {
	// overlaps: startA <= startB <= endA || startA <= endB <=endA || (startA <= startB && endA >= endB)
	if a[0] <= b[0] && a[1] >= b[0] {
		return [][]int{{a[0], max(a[1], b[1])}}
	}
	if a[0] <= b[1] && a[1] >= b[1] {
		return [][]int{{b[0], a[1]}}
	}
	if a[0] <= b[0] && a[1] >= b[1] {
		return [][]int{a}
	}
	if b[0] <= a[0] && b[1] >= a[1] {
		return [][]int{b}
	}

	return [][]int{a, b}
}

// #1 brute force
func mergeBruteForce(intervals [][]int) [][]int {
	temp := make([][]int, len(intervals))
	copy(temp, intervals)

	overlaps := make(map[int]struct{})

	tLen := len(temp)
	for i := 0; i < tLen; i++ {
		_, skip := overlaps[i]
		if skip {
			continue
		}
		cursor := i
		for j := i + 1; j < tLen; j++ {
			_, skip := overlaps[j]
			if skip {
				continue
			}
			merged := mergeTwoElements(temp[cursor], temp[j])
			if len(merged) == 1 {
				overlaps[cursor] = struct{}{}
				cursor = j
				temp[j] = merged[0]
			}
		}
	}

	ret := make([][]int, 0)
	for i, v := range temp {
		_, skip := overlaps[i]
		if skip {
			continue
		}
		ret = append(ret, v)
	}

	return ret
}

// #2 sorted list and merge two elements ordered by comparing the end of elements.end --ignore

// #3 doubly-linked table
type Node struct {
	Prev  *Node
	Next  *Node
	Value []int
}

func parseDoublyLinkedTable(intervals [][]int) *Node {
	var head *Node
	var prevNode *Node
	for _, v := range intervals {
		node := &Node{Value: v}
		if prevNode != nil {
			prevNode.Next = node
			node.Prev = prevNode
		} else {
			head = node
		}
		prevNode = node
	}

	return head
}

func mergeBySingleLinkedTable(intervals [][]int) [][]int {
	head := parseDoublyLinkedTable(intervals)

	walkNode := head.Next
	currentNode := head
	for walkNode != nil {
		cursorNode := walkNode
		for cursorNode != nil {
			merged := mergeTwoElements(currentNode.Value, cursorNode.Value)
			if len(merged) == 1 {
				if walkNode == cursorNode {
					walkNode = walkNode.Next
				}
				if head == currentNode {
					head = currentNode.Next
				} else {
					// delete currentNode
					currentNode.Next.Prev = currentNode.Prev
					currentNode.Prev.Next = currentNode.Next
				}
				cursorNode.Value = merged[0]
				currentNode = cursorNode
			}
			cursorNode = cursorNode.Next
		}
		if walkNode == nil {
			break
		}
		currentNode = walkNode
		walkNode = walkNode.Next
	}

	ret := make([][]int, 0)
	for head != nil {
		ret = append(ret, head.Value)
		head = head.Next
	}

	return ret
}

// @lc code=end
func main() {
	assert.AssertEqual1[[][]int]("[[1,3],[2,6],[8,10],[15,18]]", [][]int{{1, 6}, {8, 10}, {15, 18}}, merge, [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}})
	assert.AssertEqual1[[][]int]("[[1,4],[4,5]]", [][]int{{1, 5}}, merge, [][]int{{1, 4}, {4, 5}})
	assert.AssertEqual1[[][]int]("[[1,4],[0,5]]", [][]int{{0, 5}}, merge, [][]int{{1, 4}, {0, 5}})
	assert.AssertEqual1[[][]int]("[[1,4],[0,2],[3,5]]", [][]int{{0, 5}}, merge, [][]int{{1, 4}, {0, 2}, {3, 5}})
	assert.AssertEqual1[[][]int]("[[4,5],[2,4],[4,6],[3,4],[0,0],[1,1],[3,5],[2,2]]", [][]int{{0, 0}, {1, 1}, {2, 6}}, merge, [][]int{{4, 5}, {2, 4}, {4, 6}, {3, 4}, {0, 0}, {1, 1}, {3, 5}, {2, 2}})
	assert.AssertEqual1[[][]int]("[[2,3],[4,5],[6,7],[8,9],[1,10]]", [][]int{{1, 10}}, merge, [][]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}})
}
