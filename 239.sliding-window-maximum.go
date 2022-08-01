/*
 * @lc app=leetcode id=239 lang=golang
 *
 * [239] Sliding Window Maximum
 */
package main

import (
	"fmt"
	"sort"

	utils "reply2future.com/utils"
)

// @lc code=start
func maxSlidingWindow(nums []int, k int) []int {
	// return bruteForceMaxSlidingWindow(nums, k)
	// return binaryInsertMaxSlidingWindow(nums, k)
	return dequeMaxSlidingWindow(nums, k)
}

// #1 brute force -- time limit exceeded
func bruteForceMaxSlidingWindow(nums []int, k int) (ret []int) {
	nLen := len(nums)
	for i := 0; i+k <= nLen; i++ {
		cNums := make([]int, k)
		copy(cNums, nums[i:i+k])

		sort.Ints(cNums)
		ret = append(ret, cNums[k-1])
	}
	return
}
// #1 ====================================

// #2 binary insert -- time limit exceeded
type SlidingNode struct {
	Value int
	Prev  *SlidingNode
	Next  *SlidingNode
}

// desc stack
type SlidingStack struct {
	head   *SlidingNode
	tail   *SlidingNode
	idxMap map[int]*SlidingNode
}

func CreateSlidingStack(nums []int) (ss SlidingStack) {
	head := &SlidingNode{}
	tail := &SlidingNode{}
	head.Next = tail
	tail.Prev = head
	ss = SlidingStack{head: head, tail: tail, idxMap: make(map[int]*SlidingNode)}

	for i := 0; i < len(nums); i++ {
		ss.Unshift(i, nums[i])
	}

	return
}

func (this *SlidingStack) Shift(index int) {
	node, ok := this.idxMap[index]
	if !ok {
		panic(fmt.Sprintf("idxMap does not found:%d", index))
	}

	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
	delete(this.idxMap, index)
}

func (this *SlidingStack) Unshift(index int, value int) {
	node := &SlidingNode{Value: value}
	this.idxMap[index] = node

	var bryArray []*SlidingNode
	cNode := this.head.Next
	for cNode != this.tail {
		bryArray = append(bryArray, cNode)
		cNode = cNode.Next
	}

	// binary insert
	iNode := binaryInsertStack(bryArray, value, this.head)
	node.Next = iNode.Next
	iNode.Next.Prev = node
	iNode.Next = node
	node.Prev = iNode
}

func binaryInsertStack(bryArray []*SlidingNode, value int, head *SlidingNode) *SlidingNode {
	sLen := len(bryArray)
	mid := sLen / 2

	if sLen == 0 {
		return head
	}

	if bryArray[mid].Value > value {
		return binaryInsertStack(bryArray[mid+1:], value, bryArray[mid])
	} else if bryArray[mid].Value < value {
		return binaryInsertStack(bryArray[:mid], value, head)
	} else {
		return bryArray[mid]
	}
}

func (this SlidingStack) Max() int {
	return this.head.Next.Value
}

func binaryInsertMaxSlidingWindow(nums []int, k int) (ret []int) {
	nLen := len(nums)
	// []{value, index}
	stack := CreateSlidingStack(nums[:k])
	ret = append(ret, stack.Max())
	for i := 1; i+k <= nLen; i++ {
		index := i + k - 1
		stack.Shift(i - 1)
		stack.Unshift(index, nums[index])
		ret = append(ret, stack.Max())
	}
	return
}
// #2 ==============================================

// #3 deque
type Deque struct {
	indices []int
}

func (this *Deque) Push(index int) {
	this.indices = append(this.indices, index)
}

func (this *Deque) Pop() {
	iLen := len(this.indices)
	this.indices = this.indices[:iLen-1]
}

func (this *Deque) Shift() {
	this.indices = this.indices[1:]
}

func (this *Deque) Unshift(index int) {
	this.indices = append([]int{index}, this.indices...)
}

func (this Deque) IsEmpty() bool {
	return len(this.indices) == 0
}

func (this Deque) GetFirst() int {
	return this.indices[0]
}

func (this Deque) GetLast() int {
	iLen := len(this.indices)
	return this.indices[iLen-1]
}

// [max_value_index, except_max_max_value_index, ...]
func dequeMaxSlidingWindow(nums []int, k int) (ret []int) {
	deque := Deque{}
	
	for i := range nums {
		// unshift first element
		if !deque.IsEmpty() && i - k == deque.GetFirst() {
			deque.Shift()
		}

		// remove smaller element
		for !deque.IsEmpty() && nums[deque.GetLast()] < nums[i] {
			deque.Pop()
		}
		
		deque.Push(i)

		if i >= k - 1 {
			ret = append(ret, nums[deque.GetFirst()])
		}
	}
	return
}
// #3 =======================================

// @lc code=end
func main() {
	utils.AssertEqual2[[]int, []int, int]("example 1", []int{3, 3, 5, 5, 6, 7}, maxSlidingWindow, []int{1, 3, -1, -3, 5, 3, 6, 7}, 3)
	utils.AssertEqual2[[]int, []int, int]("example 2", []int{1}, maxSlidingWindow, []int{1}, 1)
	utils.AssertEqual2[[]int, []int, int]("example 3", []int{3, 3, 2, 5}, maxSlidingWindow, []int{1, 3, 1, 2, 0, 5}, 3)
}
