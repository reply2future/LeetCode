/*
 * @lc app=leetcode id=331 lang=golang
 *
 * [331] Verify Preorder Serialization of a Binary Tree
 */
package main

import (
	utils "reply2future.com/utils"
	"strings"
)

// @lc code=start
const CHILDREN = 2

type Stack331 []int

func (s Stack331) Len() int    { return len(s) }
func (s *Stack331) Push(i int) { *s = append(*s, i) }
func (s *Stack331) Pop() int {
	sLen := s.Len()
	ret := (*s)[sLen-1]
	*s = (*s)[:sLen-1]
	return ret
}
func (s Stack331) Peek() int {
	if s.Len() == 0 {
		panic("stack is empty")
	}
	return s[s.Len()-1]
}
func (s Stack331) IsEmpty() bool { return s.Len() == 0 }

func isValidSerialization(preorder string) bool {
	// return stackFn331(preorder)
	return simple331(preorder)
}

// #1 stack =================================================================
func stackFn331(preorder string) bool {
	if preorder == "#" {
		return true
	}

	chars := strings.Split(preorder, ",")
	stack := make(Stack331, 0)

	for i, char := range chars {
		if char == "#" {
			if stack.IsEmpty() || stack.Peek() == 0 {
				return false
			}

			for {
				stack[stack.Len()-1]--
				if stack.Peek() != 0 {
					break
				}

				stack.Pop()
				if stack.IsEmpty() {
					break
				}
			}
		} else {
			if stack.IsEmpty() && i != 0 {
				return false
			}
			stack.Push(CHILDREN)
		}
	}

	return stack.Len() == 0
}
// #1 stack =================================================================

// #2 simple way =================================================================
func simple331(preorder string) bool {
	chars := strings.Split(preorder, ",")
    diff := 1
    for _, char := range chars {
		diff--
        if diff < 0 { return false }
        if char != "#" { diff += 2 }
    }
    return diff == 0
}
// #2 simple way =================================================================
// @lc code=end
func main() {
	utils.AssertEqual1[bool, string]("example 1", true, isValidSerialization, "9,3,4,#,#,1,#,#,2,#,6,#,#")
	utils.AssertEqual1[bool, string]("example 2", false, isValidSerialization, "1,#")
	utils.AssertEqual1[bool, string]("example 3", false, isValidSerialization, "9,#,#,1")
	utils.AssertEqual1[bool, string]("example 4", true, isValidSerialization, "#")
	utils.AssertEqual1[bool, string]("example 5", false, isValidSerialization, "9,3,4,#,#,1,#,#,#,2,#,6,#,#")
}
