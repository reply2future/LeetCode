/*
 * @lc app=leetcode id=23 lang=golang
 *
 * [23] Merge k Sorted Lists
 */
package main
import (
	"fmt"
	"strconv"
)

type ListNode struct {
    Val int
    Next *ListNode
}
// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */ 
func mergeKLists(lists []*ListNode) *ListNode {
	lLen := len(lists)
	if lists == nil || lLen == 0 {
		return nil
	}

	if lLen == 1 {
		return lists[0]
	}

	if lLen != 2 {
		mid := lLen / 2
		return mergeKLists([]*ListNode{mergeKLists(lists[:mid]), mergeKLists(lists[mid:])})
	}

	var node1 *ListNode
	var node2 *ListNode
	if lists[0] == nil {
		return lists[1]
	}
	if lists[1] == nil {
		return lists[0]
	}

	if lists[0].Val < lists[1].Val {
		node1 = lists[0]
		node2 = lists[1]
	} else {
		node1 = lists[1]
		node2 = lists[0]
	}
	cNode1 := node1
	preNode1 := node1
	cNode2 := node2
	for {
		if cNode1 == nil {
			preNode1.Next = cNode2
			break
		}

		if cNode2 == nil {
			break
		}

		if cNode1.Val <= cNode2.Val {
			preNode1 = cNode1
			cNode1 = cNode1.Next
			continue
		}

		preNode1.Next = cNode2
		cNode2 = cNode2.Next
		preNode1.Next.Next = cNode1
		preNode1 = preNode1.Next 
	}
	return node1
}
// @lc code=end
func printKLists(node *ListNode) string {
	if node == nil {
		return "[]"
	}

	ret := "["
	currentNode := node
	for {
		ret += strconv.Itoa(currentNode.Val)
		if currentNode.Next == nil {
			break
		}
		ret += "->"
		currentNode = currentNode.Next 
	}
	ret += "]"

	return ret
}

func generateList(s []int) *ListNode {
	sLen := len(s)
	if sLen == 0 {
		return &ListNode{}
	}
	initNode := &ListNode{s[0], nil}
	currentNode := initNode
	for i := 1; i < len(s); i++ {
		node := &ListNode{s[i], nil}
		currentNode.Next = node
		currentNode = node
	}
	return initNode
}

func main() {
	// fmt.Println(printKLists(mergeKLists([]*ListNode{})))
	fmt.Println(printKLists(mergeKLists([]*ListNode{generateList([]int{1, 4, 5}), generateList([]int{1, 3, 4}), generateList([]int{2, 6})})))
}
