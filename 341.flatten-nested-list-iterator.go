/*
 * @lc app=leetcode id=341 lang=golang
 *
 * [341] Flatten Nested List Iterator
 */
package main

import (
	utils "reply2future.com/utils"
)

type NestedInteger struct {
	IntVal  int
	ListVal []*NestedInteger
}

// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (this NestedInteger) IsInteger() bool {
	return this.ListVal == nil
}

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (this NestedInteger) GetInteger() int {
	// if this.ListVal != nil {
	// 	return nil
	// }
	return this.IntVal
}

// Set this NestedInteger to hold a single integer.
func (this *NestedInteger) SetInteger(value int) {
	this.IntVal = value
}

// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (this *NestedInteger) Add(elem NestedInteger) {
	if (*this).ListVal == nil {
		(*this).ListVal = make([]*NestedInteger, 0)
	}

	(*this).ListVal = append((*this).ListVal, &elem)
}

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (this NestedInteger) GetList() []*NestedInteger {
	return this.ListVal
}

// @lc code=start
/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */

type NestedIterator struct {
	Stack []*NestedInteger
}

func IsEmptyNestedInteger(n NestedInteger) bool {
	if n.IsInteger() {
		return false
	}

	for _, v := range n.GetList() {
		if !IsEmptyNestedInteger(*v) {
			return false
		}
	}
	return true
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	nLen := len(nestedList)
	stack := make([]*NestedInteger, 0)
	for i := nLen - 1; i >= 0; i-- {
		if IsEmptyNestedInteger(*nestedList[i]) {
			continue
		}
		stack = append(stack, nestedList[i])
	}

	return &NestedIterator{
		Stack: stack,
	}
}

func (this *NestedIterator) Next() int {
	for {
		lastIdx := len(this.Stack) - 1
		lastElem := this.Stack[lastIdx]
		(*this).Stack = (*this).Stack[:lastIdx]

		if lastElem.IsInteger() {
			return lastElem.GetInteger()
		}

		list := lastElem.GetList()
		for i := len(list) - 1; i >= 0; i-- {
			if IsEmptyNestedInteger(*list[i]) {
				continue
			}
			(*this).Stack = append((*this).Stack, list[i])
		}
	}
}

func (this *NestedIterator) HasNext() bool {
	return len(this.Stack) > 0
}

// @lc code=end
func main() {
	n1 := &NestedInteger{
		ListVal: []*NestedInteger{
			&NestedInteger{IntVal: 1},
			&NestedInteger{IntVal: 1},
		},
	}

	n2 := &NestedInteger{IntVal: 2}

	n3 := &NestedInteger{
		ListVal: []*NestedInteger{
			&NestedInteger{IntVal: 1},
			&NestedInteger{IntVal: 1},
		},
	}

	n4 := []*NestedInteger{n1, n2, n3}
	ret := make([]int, 0)

	v := Constructor(n4)
	for i := 0; i < 5; i++ {
		ret = append(ret, v.Next())
	}
	utils.AssertEqual0[[]int]("example", []int{1, 1, 2, 1, 1}, func() []int { return ret })
}
